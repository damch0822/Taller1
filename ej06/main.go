package main

import "fmt"

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: al eliminar numeros[i] el siguiente elemento ocupa la posición i,
// pero el for hace i++ inmediatamente, saltándose ese elemento.
// Con ceros consecutivos el segundo cero queda invisible.
// Regla: por cada bloque de k ceros consecutivos sobrevive ⌊k/2⌋.

func sinCerosOriginal(numeros []int) []int {
	for i := 0; i < len(numeros); i++ {
		if numeros[i] == 0 {
			numeros = append(numeros[:i], numeros[i+1:]...)
			// i++ del for ocurre aquí → el nuevo numeros[i] no se revisa
		}
	}
	return numeros
}

// ── Parte C: versión corregida con técnica de dos índices ────────────────────
//
// filtrar recorre numeros una sola vez con dos índices: r (lectura) y w (escritura).
// Solo avanza w cuando el elemento cumple la condición; r avanza siempre.
// Sin reservar un slice nuevo: reutiliza el mismo arreglo subyacente.
func filtrar(numeros []int, cumple func(int) bool) []int {
	w := 0
	for r := 0; r < len(numeros); r++ {
		if cumple(numeros[r]) {
			numeros[w] = numeros[r]
			w++
		}
	}
	return numeros[:w]
}

func main() {
	fmt.Println("=== Parte B — código original ===")
	fmt.Println(sinCerosOriginal([]int{1, 0, 2, 0, 3}))  // [1 2 3]  ← correcto
	fmt.Println(sinCerosOriginal([]int{1, 0, 0, 2}))     // incorrecto
	fmt.Println(sinCerosOriginal([]int{0, 0, 0}))        // incorrecto

	fmt.Println("\n=== Parte C — filtrar (dos índices) ===")
	// Eliminar ceros
	fmt.Println("sin ceros {1,0,0,2}:     ", filtrar([]int{1, 0, 0, 2}, func(n int) bool { return n != 0 }))
	fmt.Println("sin ceros {0,0,0}:       ", filtrar([]int{0, 0, 0}, func(n int) bool { return n != 0 }))
	fmt.Println("sin ceros {0,1,0,0,1}:   ", filtrar([]int{0, 1, 0, 0, 1}, func(n int) bool { return n != 0 }))
	fmt.Println("sin ceros {}:            ", filtrar([]int{}, func(n int) bool { return n != 0 }))
	fmt.Println("sin ceros {0}:           ", filtrar([]int{0}, func(n int) bool { return n != 0 }))

	// Solo pares
	fmt.Println("\nsolo pares {1,2,3,4,5}:  ", filtrar([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }))

	// Mayores que umbral
	umbral := 3
	fmt.Printf("mayores que %d {1,2,3,4,5}: %v\n", umbral,
		filtrar([]int{1, 2, 3, 4, 5}, func(n int) bool { return n > umbral }))
}
