package main

import (
	"errors"
	"fmt"
)

// ── Parte B: demostración del defecto ────────────────────────────────────────
//
// DEFECTO: inscritos[:2] comparte el arreglo subyacente con inscritos.
// La capacidad de inscritos[:2] NO es 2, es 4 (la capacidad del arreglo original).
// append ve que hay espacio (len=2, cap=4) y escribe "Elsa" en la posición [2]
// del arreglo compartido, sobreescribiendo "Carla".
// Por eso inscritos queda ["Ana","Beto","Elsa","Diego"] en lugar de ["Ana","Beto","Carla","Diego"].

func demoDefecto() {
	inscritos := []string{"Ana", "Beto", "Carla", "Diego"}
	fmt.Printf("inscritos: %v  len=%d  cap=%d\n", inscritos, len(inscritos), cap(inscritos))

	sub := inscritos[:2]
	fmt.Printf("inscritos[:2]: %v  len=%d  cap=%d\n", sub, len(sub), cap(sub))

	grupoA := append(inscritos[:2], "Elsa")
	fmt.Println("grupoA   :", grupoA)
	fmt.Println("inscritos:", inscritos) // Carla fue sobreescrita
}

// ── Parte C: versiones que garantizan independencia ──────────────────────────

// copiaSegura devuelve una copia totalmente independiente usando make + copy.
func copiaSegura(s []string) []string {
	c := make([]string, len(s))
	copy(c, s)
	return c
}

// dividirGrupos divide inscritos en dos grupos independientes (sin memoria compartida).
func dividirGrupos(inscritos []string, corte int) ([]string, []string, error) {
	if corte < 0 || corte > len(inscritos) {
		return nil, nil, errors.New("corte fuera de rango")
	}
	a := make([]string, corte)
	copy(a, inscritos[:corte])
	b := make([]string, len(inscritos)-corte)
	copy(b, inscritos[corte:])
	return a, b, nil
}

// inspeccionar imprime nombre, contenido, len, cap y dirección del primer elemento.
func inspeccionar(nombre string, s []string) {
	if len(s) == 0 {
		fmt.Printf("[%s] len=%d cap=%d  (vacío)\n", nombre, len(s), cap(s))
		return
	}
	fmt.Printf("[%s] %v  len=%d  cap=%d  &[0]=%p\n", nombre, s, len(s), cap(s), &s[0])
}

func main() {
	fmt.Println("=== Parte B — defecto ===")
	demoDefecto()

	fmt.Println("\n=== Experimento con [:2:2] (tercer índice limita capacidad) ===")
	inscritos := []string{"Ana", "Beto", "Carla", "Diego"}
	grupoB := append(inscritos[:2:2], "Elsa") // cap del sub-slice = 2 → append reserva nuevo arreglo
	fmt.Println("grupoB   :", grupoB)
	fmt.Println("inscritos:", inscritos) // intacto

	fmt.Println("\n=== Parte C — dividirGrupos ===")
	original := []string{"Ana", "Beto", "Carla", "Diego", "Elsa"}
	inspeccionar("original", original)
	a, b, err := dividirGrupos(original, 3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	inspeccionar("grupoA", a)
	inspeccionar("grupoB", b)

	// Verificar independencia: modificar a no afecta original ni b
	a[0] = "MODIFICADO"
	inspeccionar("original tras modificar grupoA", original)
	inspeccionar("grupoB  tras modificar grupoA", b)

	fmt.Println("\n=== append reserva arreglo nuevo cuando cap se agota ===")
	s := make([]string, 2, 2)
	s[0], s[1] = "X", "Y"
	inspeccionar("s antes de append", s)
	s2 := append(s, "Z")
	inspeccionar("s  tras append (arreglo nuevo)", s)
	inspeccionar("s2 tras append (arreglo nuevo)", s2)
}
