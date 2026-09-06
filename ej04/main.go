package main

import (
	"errors"
	"fmt"
)

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: dentro del if se usa := en lugar de =.
//   maximo := n  →  declara una nueva variable local al bloque if.
//   La variable externa maximo nunca se actualiza, siempre vale 0.
//   Por eso la condición n > maximo se cumple incluso para 9 (9 > 0 es true),
//   y el "maximo final" imprime 0.

func demoOriginal() {
	notas := []int{12, 18, 9, 15}
	maximo := 0
	posicion := 0

	for i, n := range notas {
		if n > maximo {
			maximo := n // ← declara variable NUEVA en el bloque if (shadowing)
			posicion = i
			fmt.Println("nuevo maximo:", maximo, "en la posicion", posicion)
		}
	}
	fmt.Println("Maximo final:", maximo, "en la posicion", posicion)
}

// ── Parte C: versión corregida ────────────────────────────────────────────────
//
// maximoDe devuelve el valor máximo y su posición.
// Inicializa con el primer elemento (no con 0) para manejar correctamente
// slices con todos los valores negativos.
func maximoDe(notas []int) (int, int, error) {
	if len(notas) == 0 {
		return 0, 0, errors.New("slice vacío")
	}
	// Inicializar con el primer elemento, no con 0.
	// Si todos los valores son negativos, 0 nunca se superaría y el resultado
	// sería incorrecto (devolvería 0 como "máximo").
	maximo := notas[0]
	posicion := 0
	for i := 1; i < len(notas); i++ {
		if notas[i] > maximo {
			maximo = notas[i] // = asigna a la variable existente
			posicion = i
		}
	}
	return maximo, posicion, nil
}

// minMax recorre el slice una sola vez y devuelve mínimo y máximo.
func minMax(notas []int) (int, int, error) {
	if len(notas) == 0 {
		return 0, 0, errors.New("slice vacío")
	}
	min, max := notas[0], notas[0]
	for _, n := range notas[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max, nil
}

func main() {
	fmt.Println("=== Parte B — código original (con shadowing) ===")
	demoOriginal()

	fmt.Println("\n=== Parte C — versión corregida ===")
	casos := [][]int{
		{12, 18, 9, 15},
		{5, 3, 1},         // ordenado desc
		{-5, -3, -10},     // todos negativos (revela el problema de inicializar con 0)
		{7},               // un solo elemento
	}
	for _, notas := range casos {
		max, pos, err := maximoDe(notas)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("notas=%v  maximo=%d  posicion=%d\n", notas, max, pos)
	}

	fmt.Println("\n=== minMax (una sola pasada) ===")
	notas := []int{12, 18, 9, 15, -3}
	min, max, err := minMax(notas)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("notas=%v  min=%d  max=%d\n", notas, min, max)
	}

	// Manejo de error sin panic
	_, _, err = maximoDe([]int{})
	if err != nil {
		fmt.Println("Slice vacío detectado:", err)
	}
}
