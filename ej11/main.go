package main

import (
	"errors"
	"fmt"
)

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO en sumaRegion:
//   fórmula usada:  p[f2+1][c2+1] - p[f1][c2+1] - p[f2+1][c1] - p[f1][c1]
//   fórmula correcta: p[f2+1][c2+1] - p[f1][c2+1] - p[f2+1][c1] + p[f1][c1]
//
// Al restar la franja de arriba y la franja de la izquierda, la esquina
// superior-izquierda p[f1][c1] queda restada DOS veces.
// Hay que SUMARLA una vez para corregir.
// Cuando f1=0 o c1=0, p[f1][c1]=p[0][...]=0, por lo que el error es invisible.

func prefijos(datos [][]int) [][]int {
	n, m := len(datos), len(datos[0])
	p := make([][]int, n+1)
	for i := range p {
		p[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p[i+1][j+1] = datos[i][j] + p[i][j+1] + p[i+1][j] - p[i][j]
		}
	}
	return p
}

func sumaRegionOriginal(p [][]int, f1, c1, f2, c2 int) int {
	return p[f2+1][c2+1] - p[f1][c2+1] - p[f2+1][c1] - p[f1][c1] // BUG: - en lugar de +
}

// ── Parte C: versión corregida ────────────────────────────────────────────────

// sumaRegion aplica la fórmula de inclusión-exclusión correcta.
// Derivación:
//   P[f2+1][c2+1] = suma de todo el rectángulo (0,0)-(f2,c2)
//   P[f1][c2+1]   = suma de la franja de arriba  (0,0)-(f1-1,c2)
//   P[f2+1][c1]   = suma de la franja izquierda  (0,0)-(f2,c1-1)
//   P[f1][c1]     = esquina superior-izquierda (0,0)-(f1-1,c1-1)
//                   fue restada dos veces → debe sumarse una.
func sumaRegion(p [][]int, f1, c1, f2, c2 int) (int, error) {
	n := len(p) - 1
	m := len(p[0]) - 1
	if f1 < 0 || c1 < 0 || f2 >= n || c2 >= m {
		return 0, errors.New("coordenadas fuera de rango")
	}
	if f1 > f2 || c1 > c2 {
		return 0, errors.New("f1>f2 o c1>c2")
	}
	return p[f2+1][c2+1] - p[f1][c2+1] - p[f2+1][c1] + p[f1][c1], nil
}

// sumaRegionFuerzaBruta suma celda a celda (para validar).
func sumaRegionFuerzaBruta(datos [][]int, f1, c1, f2, c2 int) int {
	total := 0
	for i := f1; i <= f2; i++ {
		for j := c1; j <= c2; j++ {
			total += datos[i][j]
		}
	}
	return total
}

// validar compara prefijos vs fuerza bruta para TODAS las regiones posibles.
func validar(datos [][]int) bool {
	p := prefijos(datos)
	n, m := len(datos), len(datos[0])
	for f1 := 0; f1 < n; f1++ {
		for c1 := 0; c1 < m; c1++ {
			for f2 := f1; f2 < n; f2++ {
				for c2 := c1; c2 < m; c2++ {
					esperado := sumaRegionFuerzaBruta(datos, f1, c1, f2, c2)
					obtenido, err := sumaRegion(p, f1, c1, f2, c2)
					if err != nil || obtenido != esperado {
						fmt.Printf("DISCREPANCIA (%d,%d)-(%d,%d): esperado=%d obtenido=%d\n",
							f1, c1, f2, c2, esperado, obtenido)
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {
	datos := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	p := prefijos(datos)

	fmt.Println("=== Parte B — código original ===")
	fmt.Println("region (0,0)-(1,1) ->", sumaRegionOriginal(p, 0, 0, 1, 1)) // 14 correcto
	fmt.Println("region (0,0)-(3,3) ->", sumaRegionOriginal(p, 0, 0, 3, 3)) // 136 correcto (p[0][0]=0)
	fmt.Println("region (1,1)-(2,2) ->", sumaRegionOriginal(p, 1, 1, 2, 2)) // incorrecto
	fmt.Println("region (2,2)-(3,3) ->", sumaRegionOriginal(p, 2, 2, 3, 3)) // incorrecto

	fmt.Println("\n=== Parte C — versión corregida ===")
	consultas := [][4]int{{0, 0, 1, 1}, {0, 0, 3, 3}, {1, 1, 2, 2}, {2, 2, 3, 3}}
	for _, q := range consultas {
		f1, c1, f2, c2 := q[0], q[1], q[2], q[3]
		res, _ := sumaRegion(p, f1, c1, f2, c2)
		bf := sumaRegionFuerzaBruta(datos, f1, c1, f2, c2)
		ok := "✓"
		if res != bf {
			ok = "✗"
		}
		fmt.Printf("(%d,%d)-(%d,%d)  prefijo=%d  bruteforce=%d  %s\n", f1, c1, f2, c2, res, bf, ok)
	}

	fmt.Println("\n=== Validación completa de todas las regiones ===")
	fmt.Println("Todas las regiones correctas:", validar(datos))

	fmt.Println("\n=== Error en coordenadas fuera de rango ===")
	_, err := sumaRegion(p, -1, 0, 2, 2)
	fmt.Println("Error esperado:", err)
}
