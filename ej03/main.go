package main

import (
	"errors"
	"fmt"
	"math"
)

// ── Parte B: código original con los dos defectos señalados ──────────────────
//
// DEFECTO 1 (pesosValidos): compara suma == 1.0 con ==.
//   0.30 + 0.35 + 0.35 en float64 no es exactamente 1.0 (representación binaria).
//
// DEFECTO 2 (promedioPonderado): float64(suma / totalPeso) hace la división
//   primero en enteros (trunca) y convierte el entero resultante a float64.
//   Debe ser float64(suma) / float64(totalPeso).

func pesosValidosOriginal(pesos []float64) bool {
	suma := 0.0
	for _, p := range pesos {
		suma += p
	}
	return suma == 1.0 // comparación exacta: falla con 0.30+0.35+0.35
}

func promedioPonderadoOriginal(notas []int, pesos []int) float64 {
	suma := 0
	totalPeso := 0
	for i := range notas {
		suma += notas[i] * pesos[i]
		totalPeso += pesos[i]
	}
	return float64(suma / totalPeso) // división entera antes de convertir
}

// ── Parte C: versiones corregidas ────────────────────────────────────────────

// casiIgual compara dos float64 con una tolerancia epsilon.
// Epsilon = 1e-9 es adecuado para pesos en [0,1]: la acumulación de errores
// de representación binaria para 2-3 sumandos en ese rango es del orden de 1e-16,
// mucho menor que 1e-9, por lo que no hay falsos positivos ni falsos negativos.
func casiIgual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

// promedioPonderado trabaja en float64 desde el inicio y valida entradas.
func promedioPonderado(notas []float64, pesos []float64) (float64, error) {
	if len(notas) == 0 || len(pesos) == 0 {
		return 0, errors.New("slices vacíos")
	}
	if len(notas) != len(pesos) {
		return 0, errors.New("longitudes distintas")
	}
	sumaPesos := 0.0
	for _, p := range pesos {
		sumaPesos += p
	}
	if !casiIgual(sumaPesos, 1.0, 1e-9) {
		return 0, fmt.Errorf("los pesos suman %.20f, no 1", sumaPesos)
	}
	resultado := 0.0
	for i := range notas {
		resultado += notas[i] * pesos[i]
	}
	return resultado, nil
}

// distribucion cuenta cuántas notas caen en cada rango usando switch sin expresión.
func distribucion(notas []float64) map[string]int {
	cont := map[string]int{
		"0-10":  0,
		"11-14": 0,
		"15-17": 0,
		"18-20": 0,
	}
	for _, n := range notas {
		switch {
		case n >= 18:
			cont["18-20"]++
		case n >= 15:
			cont["15-17"]++
		case n >= 11:
			cont["11-14"]++
		default:
			cont["0-10"]++
		}
	}
	return cont
}

func main() {
	fmt.Println("=== Parte B — código original ===")
	pesos := []float64{0.30, 0.35, 0.35}
	fmt.Println("Pesos validos (original):", pesosValidosOriginal(pesos))
	fmt.Printf("Suma real con %%.20f: %.20f\n", 0.30+0.35+0.35)

	notas := []int{14, 16, 11}
	pesosPct := []int{30, 35, 35}
	fmt.Printf("Promedio (original): %.2f\n", promedioPonderadoOriginal(notas, pesosPct))

	fmt.Println("\n=== Parte B — casos de verificación para pesosValidos ===")
	casos := [][]float64{
		{0.2, 0.3, 0.5},
		{0.1, 0.2, 0.7},
		{0.30, 0.35, 0.35},
	}
	for _, c := range casos {
		suma := 0.0
		for _, p := range c {
			suma += p
		}
		fmt.Printf("pesos=%v  suma=%.20f  ==1.0? %v\n", c, suma, suma == 1.0)
	}

	fmt.Println("\n=== Parte C — versión corregida ===")
	notasF := []float64{14, 16, 11}
	pesosF := []float64{0.30, 0.35, 0.35}
	res, err := promedioPonderado(notasF, pesosF)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Promedio ponderado corregido: %.2f\n", res)
	}

	fmt.Println("\nDistribución:", distribucion([]float64{8, 12, 15, 19, 10, 17, 20}))
}
