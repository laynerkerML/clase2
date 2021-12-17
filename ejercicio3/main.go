package main

import "fmt"

func main() {
	SalarioTotal := calculoSalario(40, "A")
	fmt.Println("Mi Sueldo es: ", SalarioTotal)
}

func calculoSalario(horas int, catego string) (salario float64) {
	salarioXHora, porcentaje := categorizarcion(catego)
	salario = float64(horas) * salarioXHora
	if catego == "C" {
		return salario
	}
	bono := salario * porcentaje
	salario += bono
	return salario
}

func categorizarcion(catego string) (salarioXHora float64, porcentaje float64) {
	salarioXHora = 0.00
	porcentaje = 0.00
	switch catego {
	case "C":
		salarioXHora = 1000.00
	case "B":
		salarioXHora = 1500.00
		porcentaje = 0.2
	case "A":
		salarioXHora = 3000.00
		porcentaje = 0.5
	}
	return salarioXHora, porcentaje
}
