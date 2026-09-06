package main

import (
	"fmt"
	"math/bits"
)

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: cuando ordenados[mid] < objetivo se hace lo = mid en vez de lo = mid+1.
// Si lo y hi son consecutivos (ej. lo=3, hi=4) entonces mid=3=lo.
// La nueva asignación lo=mid=3 no reduce el intervalo → ciclo infinito.
//
// NOTA: NO ejecutar buscarOriginal con un valor ausente (se cuelga).

func buscarOriginal(ordenados []int, objetivo int) int {
	lo, hi := 0, len(ordenados)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		switch {
		case ordenados[mid] == objetivo:
			return mid
		case ordenados[mid] < objetivo:
			lo = mid // BUG: debería ser lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

// ── Parte C: versión corregida ────────────────────────────────────────────────
//
// Invariante: si objetivo está en ordenados, está en ordenados[lo..hi].
// En cada iteración el intervalo se reduce estrictamente:
//   - si vamos a la derecha: lo = mid+1  → lo crece, intervalo disminuye.
//   - si vamos a la izquierda: hi = mid-1 → hi decrece, intervalo disminuye.
// La cantidad hi-lo+1 decrece en al menos 1 cada vuelta → termina.

func buscar(ordenados []int, objetivo int) (idx int, iteraciones int) {
	lo, hi := 0, len(ordenados)-1
	// Invariante: objetivo está en ordenados[lo..hi] o no está.
	for lo <= hi {
		iteraciones++
		mid := lo + (hi-lo)/2 // evita desbordamiento aunque en Go int es 64 bits
		switch {
		case ordenados[mid] == objetivo:
			return mid, iteraciones
		case ordenados[mid] < objetivo:
			lo = mid + 1 // CORRECCIÓN: descartamos mid
		default:
			hi = mid - 1
		}
	}
	return -1, iteraciones
}

// limiteInferior devuelve el índice del primer elemento >= objetivo,
// o len(ordenados) si todos son menores. Funciona con elementos repetidos.
func limiteInferior(ordenados []int, objetivo int) int {
	lo, hi := 0, len(ordenados)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if ordenados[mid] < objetivo {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// contarOcurrencias cuenta en O(log n) cuántas veces aparece objetivo.
func contarOcurrencias(ordenados []int, objetivo int) int {
	ini := limiteInferior(ordenados, objetivo)
	fin := limiteInferior(ordenados, objetivo+1)
	return fin - ini
}

func main() {
	fmt.Println("=== Parte B — código original (solo casos que terminan) ===")
	datos := []int{2, 5, 8, 12, 20}
	fmt.Println("buscar 12 ->", buscarOriginal(datos, 12)) // encontrado: 3
	fmt.Println("buscar 2  ->", buscarOriginal(datos, 2))  // encontrado: 0
	// buscarOriginal(datos, 15) se colgaría → NO lo ejecutamos

	fmt.Println("\n=== Parte C — versión corregida ===")
	casos := []int{2, 5, 8, 12, 20, 15, 1, 25, 7}
	for _, obj := range casos {
		idx, iters := buscar(datos, obj)
		fmt.Printf("buscar %-3d -> idx=%-2d  iteraciones=%d\n", obj, idx, iters)
	}

	fmt.Println("\n=== Verificación: iteraciones <= log2(n)+1 ===")
	for _, n := range []int{10, 100, 1000, 10000} {
		// construir slice ordenado [0, 2, 4, ..., 2*(n-1)]
		slice := make([]int, n)
		for i := range slice {
			slice[i] = i * 2
		}
		limite := bits.Len(uint(n)) // ⌊log2(n)⌋ + 1
		maxIter := 0
		for obj := -1; obj <= 2*n+1; obj++ {
			_, it := buscar(slice, obj)
			if it > maxIter {
				maxIter = it
			}
		}
		ok := "✓"
		if maxIter > limite {
			ok = "✗ SUPERA"
		}
		fmt.Printf("n=%-6d  limite=%d  maxIter=%d  %s\n", n, limite, maxIter, ok)
	}

	fmt.Println("\n=== limiteInferior y contarOcurrencias ===")
	rep := []int{1, 2, 2, 3, 3, 3, 4}
	for _, v := range []int{0, 2, 3, 5} {
		fmt.Printf("valor=%-2d  limiteInferior=%d  ocurrencias=%d\n",
			v, limiteInferior(rep, v), contarOcurrencias(rep, v))
	}
}
