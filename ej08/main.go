package main

import (
	"errors"
	"fmt"
)

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: el ciclo de transposición recorre j de 0 a n-1 para cada i.
// Esto intercambia el par (i,j)/(j,i) DOS veces: primero cuando i<j y luego
// cuando i>j (los roles se invierten). El resultado neto es la identidad:
// la matriz queda igual. Solo la inversión de filas (segunda etapa) se aplica,
// produciendo una rotación incorrecta.

func rotarOriginal(m [][]int) {
	n := len(m)
	// Transposición CON DEFECTO: j va de 0 a n (recorre toda la fila)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	// Invertir filas (correcto)
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			m[i][j], m[i][k] = m[i][k], m[i][j]
		}
	}
}

// ── Parte C: versiones corregidas ────────────────────────────────────────────

// rotar gira la matriz 90° en sentido horario, in situ.
// Etapa 1: transponer recorriendo solo j > i (triángulo superior) para no deshacer el intercambio.
// Etapa 2: invertir cada fila horizontalmente.
func rotar(m [][]int) error {
	n := len(m)
	if n == 0 {
		return errors.New("matriz vacía")
	}
	for _, fila := range m {
		if len(fila) != n {
			return errors.New("matriz no cuadrada")
		}
	}
	// Transposición CORRECTA: j empieza en i+1 → cada par se intercambia UNA sola vez.
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	// Invertir filas
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			m[i][j], m[i][k] = m[i][k], m[i][j]
		}
	}
	return nil
}

// rotarAntihorario gira 90° en sentido antihorario.
// Para rotar antihorario: primero invertir filas, luego transponer.
// El orden importa: transponer+invertir y invertir+transponer son rotaciones distintas.
func rotarAntihorario(m [][]int) error {
	n := len(m)
	if n == 0 {
		return errors.New("matriz vacía")
	}
	for _, fila := range m {
		if len(fila) != n {
			return errors.New("matriz no cuadrada")
		}
	}
	// Primero invertir filas
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			m[i][j], m[i][k] = m[i][k], m[i][j]
		}
	}
	// Luego transponer
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	return nil
}

// esSimetrica comprueba que m[i][j] == m[j][i] para i < j.
func esSimetrica(m [][]int) bool {
	n := len(m)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if m[i][j] != m[j][i] {
				return false
			}
		}
	}
	return true
}

// copiarMatriz devuelve una copia independiente.
func copiarMatriz(m [][]int) [][]int {
	c := make([][]int, len(m))
	for i := range m {
		c[i] = make([]int, len(m[i]))
		copy(c[i], m[i])
	}
	return c
}

// igualesMatriz compara dos matrices elemento a elemento.
func igualesMatriz(a, b [][]int) bool {
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

// verificarRotacion aplica rotar cuatro veces y comprueba que regresa al estado inicial.
func verificarRotacion(m [][]int) bool {
	original := copiarMatriz(m)
	copia := copiarMatriz(m)
	for i := 0; i < 4; i++ {
		if err := rotar(copia); err != nil {
			return false
		}
	}
	return igualesMatriz(original, copia)
}

func imprimir(m [][]int) {
	for _, fila := range m {
		fmt.Println(fila)
	}
	fmt.Println("---")
}

func main() {
	fmt.Println("=== Parte B — código original ===")
	sim := [][]int{{1, 2}, {2, 1}}
	rotarOriginal(sim)
	imprimir(sim)

	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotarOriginal(m)
	imprimir(m)

	fmt.Println("=== Parte C — versión corregida ===")
	m2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotar(m2)
	imprimir(m2) // [[7 4 1] [8 5 2] [9 6 3]]

	fmt.Println("=== verificarRotacion (4 vueltas = identidad) ===")
	casos := [][][]int{
		{{1, 2}, {3, 4}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
		{{1}},
	}
	for _, c := range casos {
		fmt.Printf("matriz %dx%d: 4 rotaciones = identidad? %v\n", len(c), len(c[0]), verificarRotacion(c))
	}

	fmt.Println("\n=== esSimetrica ===")
	fmt.Println("[[1,2],[2,1]] simétrica?", esSimetrica([][]int{{1, 2}, {2, 1}}))
	fmt.Println("[[1,2],[3,1]] simétrica?", esSimetrica([][]int{{1, 2}, {3, 1}}))
}
