package main

import "fmt"

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: dentro del cuerpo del for se hace i++ cuando i es impar.
// Luego la cláusula de incremento del for hace otro i++.
// Para n impar (ej. n=9), cuando i llega a 9 (impar), se hace i++ → i=10,
// se suma 10 (fuera del rango [1,9]) y luego el for hace i++ → i=11 > 9, termina.
// Resultado incorrecto: suma 10 en lugar de no sumar nada.

func sumParesOriginal(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			i++ // saltamos el impar (doble incremento con el del for)
		}
		total += i
	}
	return total
}

// ── Parte C: versión corregida ────────────────────────────────────────────────
//
// La variable de control sólo cambia en la cláusula del for (i += 2).
// Empezamos en 2 (primer par) y avanzamos de 2 en 2.
//
// Invariante (comentario encima del for):
//   Al inicio de cada iteración, total == suma de todos los pares en [2, i-2]
//   (es decir, ya procesamos todos los pares menores a i).
//
// Terminación: la cantidad (n - i + 2) decrece estrictamente en cada vuelta
// porque i aumenta 2, y como i está acotado por n, el ciclo termina.

func sumaPares(n int) int {
	total := 0
	// Invariante: total == suma de pares en [2, i-2] al inicio de cada iteración.
	for i := 2; i <= n; i += 2 {
		total += i
	}
	return total
}

// sumaParesFormula calcula la suma sin ciclo.
// La suma de los k primeros pares 2+4+...+2k = k*(k+1).
// k = cantidad de pares en [1,n] = n/2 (división entera).
func sumaParesFormula(n int) int {
	k := n / 2
	return k * (k + 1)
}

func main() {
	fmt.Println("=== Parte B — código original ===")
	fmt.Println("n=10 ->", sumParesOriginal(10)) // correcto: 30
	fmt.Println("n=9  ->", sumParesOriginal(9))  // incorrecto
	fmt.Println("n=1  ->", sumParesOriginal(1))  // incorrecto

	fmt.Println("\n=== Parte C — versión corregida ===")
	for _, n := range []int{0, 1, 2, 9, 10, 11} {
		iter := sumaPares(n)
		form := sumaParesFormula(n)
		match := "✓"
		if iter != form {
			match = "✗ DIFIEREN"
		}
		fmt.Printf("n=%-3d  iterativa=%d  formula=%d  %s\n", n, iter, form, match)
	}

	fmt.Println("\n=== Verificación fórmula vs iterativa n=0..100 ===")
	ok := true
	for n := 0; n <= 100; n++ {
		if sumaPares(n) != sumaParesFormula(n) {
			fmt.Printf("DIVERGEN en n=%d\n", n)
			ok = false
		}
	}
	if ok {
		fmt.Println("Todas coinciden.")
	}
}
