package main

import "fmt"

// ── Reglas del Juego de la Vida ───────────────────────────────────────────────
// Viva con 2 o 3 vecinos → sobrevive. Viva con otro nro → muere.
// Muerta con exactamente 3 vecinos → nace. Muerta, otro nro → sigue muerta.

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: siguiente() modifica el tablero IN SITU.
// Cuando evalúa la celda (i,j) en la iteración actual, algunas celdas ya
// fueron actualizadas en este mismo recorrido, contaminando el conteo de vecinos.
// La primera celda "contaminada" depende del patrón, pero en el blinker
// ocurre al evaluar (1,2): la celda (2,1) ya se modificó en el mismo recorrido.

func vecinosOriginal(tab [][]bool, f, c int) int {
	total := 0
	for df := -1; df <= 1; df++ {
		for dc := -1; dc <= 1; dc++ {
			if df == 0 && dc == 0 {
				continue
			}
			i, j := f+df, c+dc
			dentro := i >= 0 && i < len(tab) && j >= 0 && j < len(tab[0])
			if dentro && tab[i][j] {
				total++
			}
		}
	}
	return total
}

func siguienteOriginal(tab [][]bool) {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[0]); j++ {
			v := vecinosOriginal(tab, i, j)
			if tab[i][j] && (v < 2 || v > 3) {
				tab[i][j] = false // modifica el estado actual → contamina vecinos futuros
			}
			if !tab[i][j] && v == 3 {
				tab[i][j] = true
			}
		}
	}
}

// ── Parte C: versión corregida ────────────────────────────────────────────────

// vecinos cuenta los vecinos en una cuadrícula con bordes fijos.
func vecinos(tab [][]bool, f, c int) int {
	total := 0
	for df := -1; df <= 1; df++ {
		for dc := -1; dc <= 1; dc++ {
			if df == 0 && dc == 0 {
				continue
			}
			i, j := f+df, c+dc
			dentro := i >= 0 && i < len(tab) && j >= 0 && j < len(tab[0])
			if dentro && tab[i][j] {
				total++
			}
		}
	}
	return total
}

// vecinosToroidal cuenta vecinos en cuadrícula toroidal (bordes conectados).
// Usa aritmética modular para envolver los índices.
func vecinosToroidal(tab [][]bool, f, c int) int {
	filas := len(tab)
	cols := len(tab[0])
	total := 0
	for df := -1; df <= 1; df++ {
		for dc := -1; dc <= 1; dc++ {
			if df == 0 && dc == 0 {
				continue
			}
			// ((f+df) % filas + filas) % filas evita módulo negativo en Go
			i := ((f + df) % filas + filas) % filas
			j := ((c + dc) % cols + cols) % cols
			if tab[i][j] {
				total++
			}
		}
	}
	return total
}

// siguiente devuelve un tablero NUEVO con las reglas aplicadas simultáneamente.
// El original NO se modifica — eso garantiza que todos los vecinos se calculan
// sobre el estado anterior, no sobre el estado parcialmente actualizado.
func siguiente(tab [][]bool) [][]bool {
	filas := len(tab)
	cols := len(tab[0])
	nuevo := make([][]bool, filas)
	for i := range nuevo {
		nuevo[i] = make([]bool, cols)
	}
	for i := 0; i < filas; i++ {
		for j := 0; j < cols; j++ {
			v := vecinos(tab, i, j)
			switch {
			case tab[i][j] && (v == 2 || v == 3):
				nuevo[i][j] = true // sobrevive
			case !tab[i][j] && v == 3:
				nuevo[i][j] = true // nace
			// en otro caso, nuevo[i][j] ya es false
			}
		}
	}
	return nuevo
}

// iguales compara dos tableros celda a celda.
func iguales(a, b [][]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// copiarTablero devuelve una copia independiente.
func copiarTablero(tab [][]bool) [][]bool {
	c := make([][]bool, len(tab))
	for i := range tab {
		c[i] = make([]bool, len(tab[i]))
		copy(c[i], tab[i])
	}
	return c
}

// simular ejecuta N generaciones y devuelve la primera generación que repite
// un estado anterior. Devuelve -1 si no se detecta ciclo.
func simular(tab [][]bool, generaciones int) int {
	historial := [][]bool{}
	// guardamos estado inicial como cadena de bits plana para comparar
	actual := copiarTablero(tab)
	historial = append(historial, flatten(actual))
	for g := 1; g <= generaciones; g++ {
		actual = siguiente(actual)
		est := flatten(actual)
		for k, prev := range historial {
			if igualesBool(est, prev) {
				return g - k // longitud del ciclo detectado en generación g
			}
		}
		historial = append(historial, est)
	}
	return -1
}

// flatten convierte [][]bool en []bool para comparación sencilla.
func flatten(tab [][]bool) []bool {
	var out []bool
	for _, fila := range tab {
		out = append(out, fila...)
	}
	return out
}

func igualesBool(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// imprimirTablero imprime el tablero con '.' y '#'.
func imprimirTablero(tab [][]bool) {
	for _, fila := range tab {
		for _, c := range fila {
			if c {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// nuevoTablero crea un tablero de filas x cols (todo false).
func nuevoTablero(filas, cols int) [][]bool {
	t := make([][]bool, filas)
	for i := range t {
		t[i] = make([]bool, cols)
	}
	return t
}

func main() {
	// Blinker horizontal: (2,1),(2,2),(2,3) en tablero 5x5
	tab := nuevoTablero(5, 5)
	tab[2][1] = true
	tab[2][2] = true
	tab[2][3] = true

	fmt.Println("=== Parte B — código original (in situ) ===")
	orig := copiarTablero(tab)
	fmt.Println("Generación 0:")
	imprimirTablero(orig)
	siguienteOriginal(orig)
	fmt.Println("Generación 1 (in situ — debería ser vertical):")
	imprimirTablero(orig)
	siguienteOriginal(orig)
	fmt.Println("Generación 2 (in situ — debería volver a horizontal):")
	imprimirTablero(orig)

	fmt.Println("=== Parte C — versión corregida ===")
	actual := copiarTablero(tab)
	fmt.Println("Generación 0:")
	imprimirTablero(actual)
	actual = siguiente(actual)
	fmt.Println("Generación 1 (debe ser vertical):")
	imprimirTablero(actual)
	actual = siguiente(actual)
	fmt.Println("Generación 2 (debe volver a horizontal):")
	imprimirTablero(actual)

	fmt.Println("=== Verificación: bloque 2x2 (stable) ===")
	bloque := nuevoTablero(4, 4)
	bloque[1][1] = true
	bloque[1][2] = true
	bloque[2][1] = true
	bloque[2][2] = true
	bloqueSig := siguiente(bloque)
	fmt.Println("Bloque 2x2 no cambia:", iguales(bloque, bloqueSig))

	fmt.Println("\n=== Detección de ciclo (simular) ===")
	blinker := nuevoTablero(5, 5)
	blinker[2][1] = true
	blinker[2][2] = true
	blinker[2][3] = true
	periodo := simular(blinker, 10)
	fmt.Println("Periodo del blinker:", periodo) // debe ser 2
}
