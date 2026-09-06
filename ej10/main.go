package main

import (
	"errors"
	"fmt"
)

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: `suma` se declara una sola vez ANTES de los tres ciclos.
// Debería reiniciarse a 0.0 al inicio de cada celda (i,j).
// La celda (0,0) es correcta porque suma empieza en 0.
// Para (0,1), suma YA vale 19 (de la celda anterior) y se le suman 22 más → 41.
// El valor esperado era 22. El 41 viene de 19 (c[0][0]) + 22 (acumulado nuevo).

func multiplicarOriginal(a, b [][]float64) [][]float64 {
	n := len(a)
	c := make([][]float64, n)
	for i := range c {
		c[i] = make([]float64, n)
	}

	suma := 0.0 // BUG: debe estar dentro del ciclo j
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				suma += a[i][k] * b[k][j]
			}
			c[i][j] = suma
		}
	}
	return c
}

// ── Parte C: versión corregida IJK ───────────────────────────────────────────

// multiplicar calcula el producto matricial. Devuelve error si las dimensiones
// son incompatibles (no asume que ambas son cuadradas).
func multiplicar(a, b [][]float64) ([][]float64, error) {
	filasA := len(a)
	if filasA == 0 {
		return nil, errors.New("matriz a vacía")
	}
	colsA := len(a[0])
	filasB := len(b)
	if filasB == 0 {
		return nil, errors.New("matriz b vacía")
	}
	colsB := len(b[0])
	if colsA != filasB {
		return nil, fmt.Errorf("dimensiones incompatibles: a tiene %d columnas, b tiene %d filas", colsA, filasB)
	}

	c := make([][]float64, filasA)
	for i := range c {
		c[i] = make([]float64, colsB)
	}

	for i := 0; i < filasA; i++ {
		for j := 0; j < colsB; j++ {
			suma := 0.0 // CORRECCIÓN: reiniciar por celda
			for k := 0; k < colsA; k++ {
				suma += a[i][k] * b[k][j]
			}
			c[i][j] = suma
		}
	}
	return c, nil
}

// multiplicarIKJ — orden i,k,j: para cada fila i y columna k de a,
// escribe directamente en c[i][j] sin acumulador separado.
// Esto favorece la localidad de caché (c[i] y b[k] se recorren en orden).
func multiplicarIKJ(a, b [][]float64) ([][]float64, error) {
	filasA := len(a)
	if filasA == 0 {
		return nil, errors.New("matriz a vacía")
	}
	colsA := len(a[0])
	filasB := len(b)
	if filasB == 0 {
		return nil, errors.New("matriz b vacía")
	}
	colsB := len(b[0])
	if colsA != filasB {
		return nil, fmt.Errorf("dimensiones incompatibles")
	}

	c := make([][]float64, filasA)
	for i := range c {
		c[i] = make([]float64, colsB)
	}

	for i := 0; i < filasA; i++ {
		for k := 0; k < colsA; k++ {
			aik := a[i][k]
			for j := 0; j < colsB; j++ {
				c[i][j] += aik * b[k][j]
			}
		}
	}
	return c, nil
}

// igualesMatriz compara dos matrices elemento a elemento.
func igualesMatriz(a, b [][]float64) bool {
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

func imprimir(m [][]float64) {
	for _, fila := range m {
		fmt.Println(fila)
	}
}

func main() {
	a := [][]float64{{1, 2}, {3, 4}}
	b := [][]float64{{5, 6}, {7, 8}}

	fmt.Println("=== Parte B — código original ===")
	imprimir(multiplicarOriginal(a, b)) // incorrecto

	fmt.Println("\n=== Parte C — multiplicar corregido (IJK) ===")
	c1, _ := multiplicar(a, b)
	imprimir(c1) // [[19 22] [43 50]]

	fmt.Println("\n=== Parte C — multiplicarIKJ ===")
	c2, _ := multiplicarIKJ(a, b)
	imprimir(c2)

	fmt.Println("\nAmbas producen el mismo resultado?", igualesMatriz(c1, c2))

	fmt.Println("\n=== Verificaciones adicionales ===")
	// Identidad
	id := [][]float64{{1, 0}, {0, 1}}
	ai, _ := multiplicar(a, id)
	fmt.Println("A x I == A?", igualesMatriz(ai, a))

	// Dimensiones incompatibles
	_, err := multiplicar([][]float64{{1, 2}}, [][]float64{{1}, {2}, {3}})
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	// Matrices 3x3
	a3 := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	b3 := [][]float64{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
	r1, _ := multiplicar(a3, b3)
	r2, _ := multiplicarIKJ(a3, b3)
	fmt.Println("3x3 IJK == IKJ?", igualesMatriz(r1, r2))
}
