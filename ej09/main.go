package main

import (
	"errors"
	"fmt"
)

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: Tabla es [3][4]float64, un ARRAY (valor, no referencia).
// Al pasarlo a normalizar, Go copia TODOS los 12 float64.
// Los cambios dentro de la función ocurren sobre la copia; el original
// en main nunca se modifica. El compilador no avisa porque la función
// es sintácticamente válida.

type Tabla [3][4]float64

func normalizarOriginal(t Tabla) {
	for i := 0; i < len(t); i++ {
		maximo := t[i][0]
		for j := 1; j < len(t[i]); j++ {
			if t[i][j] > maximo {
				maximo = t[i][j]
			}
		}
		for j := 0; j < len(t[i]); j++ {
			t[i][j] = t[i][j] / maximo
		}
	}
	fmt.Println("  [dentro de normalizar] fila 0:", t[0]) // valores modificados
}

// ── Parte C: solución A — puntero al array ────────────────────────────────────

// normalizarPtr recibe un puntero: no copia el array.
// Devuelve error si alguna fila tiene máximo == 0.
func normalizarPtr(t *Tabla) error {
	for i := 0; i < len(t); i++ {
		maximo := t[i][0]
		for j := 1; j < len(t[i]); j++ {
			if t[i][j] > maximo {
				maximo = t[i][j]
			}
		}
		if maximo == 0 {
			return fmt.Errorf("fila %d tiene máximo 0, no se puede normalizar", i)
		}
		for j := 0; j < len(t[i]); j++ {
			t[i][j] /= maximo
		}
	}
	return nil
}

// ── Parte C: solución B — slice de slices ────────────────────────────────────

// normalizarSlice recibe [][]float64.
// Un slice es una cabecera (puntero + len + cap) que se copia, pero el puntero
// apunta al arreglo subyacente original → los cambios se ven en el llamador.
// Por eso NO hace falta el puntero.
func normalizarSlice(t [][]float64) error {
	for i := range t {
		maximo := t[i][0]
		for j := 1; j < len(t[i]); j++ {
			if t[i][j] > maximo {
				maximo = t[i][j]
			}
		}
		if maximo == 0 {
			return errors.New("fila con máximo 0")
		}
		for j := range t[i] {
			t[i][j] /= maximo
		}
	}
	return nil
}

func main() {
	fmt.Println("=== Parte B — código original (sin efecto en main) ===")
	t := Tabla{{2, 4, 6, 8}, {1, 2, 3, 4}, {5, 10, 15, 20}}
	fmt.Println("  [antes de llamar] fila 0:", t[0])
	normalizarOriginal(t)
	fmt.Println("  [después de llamar] fila 0:", t[0]) // sin cambios

	fmt.Println("\n=== Parte C — normalizarPtr (puntero) ===")
	t2 := Tabla{{2, 4, 6, 8}, {1, 2, 3, 4}, {5, 10, 15, 20}}
	if err := normalizarPtr(&t2); err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, fila := range t2 {
			fmt.Println(fila)
		}
	}

	fmt.Println("\n=== Parte C — normalizarSlice (slice de slices) ===")
	s := [][]float64{
		{2, 4, 6, 8},
		{1, 2, 3, 4},
		{5, 10, 15, 20},
	}
	if err := normalizarSlice(s); err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, fila := range s {
			fmt.Println(fila)
		}
	}

	fmt.Println("\n=== Cuántos bytes se copian ===")
	// Tabla [3][4]float64 = 12 * 8 bytes = 96 bytes por valor
	// *Tabla = 8 bytes (puntero de 64 bits)
	// [][]float64 cabecera = 24 bytes (puntero + len + cap); datos sin copiar
	fmt.Println("Tabla por valor:  96 bytes copiados (12 float64)")
	fmt.Println("*Tabla (puntero):  8 bytes copiados")
	fmt.Println("[][]float64 cabecera: 24 bytes (el arreglo subyacente no se copia)")
}
