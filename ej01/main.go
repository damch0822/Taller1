package main

import (
	"errors"
	"fmt"
)

// ── Parte B: código original con el defecto señalado ─────────────────────────
//
// DEFECTO: la cadena if/else if evalúa >= 11 primero.
// Cualquier nota >= 11 ya entra ahí y NUNCA llega a la rama >= 16 ni >= 18.
// DESTACADO y EXCELENTE son código muerto.

func categoriaOriginal(nota int) string {
	if nota < 0 || nota > 20 {
		return "INVALIDA"
	}
	if nota >= 11 {
		return "APROBADO" // atrapa todo >= 11; las ramas siguientes son inalcanzables
	} else if nota >= 16 {
		return "DESTACADO" // nunca se evalúa
	} else if nota >= 18 {
		return "EXCELENTE" // nunca se evalúa
	}
	return "DESAPROBADO"
}

// ── Parte C: versión corregida ────────────────────────────────────────────────
//
// Reglas: 0-10 → DESAPROBADO, 11-15 → APROBADO, 16-17 → DESTACADO, 18-20 → EXCELENTE
// Las condiciones se ordenan de la más restrictiva a la más general dentro del switch.

func categoria(nota int) (string, error) {
	if nota < 0 || nota > 20 {
		return "", errors.New("nota fuera del rango 0-20")
	}
	switch {
	case nota >= 18:
		return "EXCELENTE", nil
	case nota >= 16:
		return "DESTACADO", nil
	case nota >= 11:
		return "APROBADO", nil
	default:
		return "DESAPROBADO", nil
	}
}

// tablaDeCobertura imprime las 21 notas válidas con su categoría alineada.
func tablaDeCobertura() {
	fmt.Printf("%-5s  %s\n", "Nota", "Categoria")
	fmt.Println("-----  -----------")
	for n := 0; n <= 20; n++ {
		cat, _ := categoria(n)
		fmt.Printf("%-5d  %s\n", n, cat)
	}
}

// verificarCobertura comprueba que ninguna nota de 0-20 devuelva cadena vacía
// y que las cuatro categorías aparezcan al menos una vez.
func verificarCobertura() bool {
	encontradas := map[string]bool{}
	for n := 0; n <= 20; n++ {
		cat, err := categoria(n)
		if err != nil || cat == "" {
			return false
		}
		encontradas[cat] = true
	}
	requeridas := []string{"DESAPROBADO", "APROBADO", "DESTACADO", "EXCELENTE"}
	for _, r := range requeridas {
		if !encontradas[r] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== Parte B — código original (5 notas de prueba) ===")
	for _, n := range []int{8, 11, 15, 17, 19} {
		fmt.Printf("%2d -> %s\n", n, categoriaOriginal(n))
	}

	fmt.Println("\n=== Parte C — tabla de cobertura (código corregido) ===")
	tablaDeCobertura()

	fmt.Println("\n=== Parte C — verificar cobertura ===")
	fmt.Println("Cobertura correcta:", verificarCobertura())
}
