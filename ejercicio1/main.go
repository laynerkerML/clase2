package main

import "fmt"

func main() {
	salerio := 40000.00
	calculoDescuento := calcularImpuesto(salerio)
	salarioPagar := salerio - calculoDescuento
	fmt.Println("Salario total: ", salarioPagar)
	fmt.Println("Impuesto: ", calculoDescuento)
}

func calcularImpuesto(salario float64) float64 {
	porcentaje := 0.00

	if salario > 50000.00 {
		porcentaje = 0.17
	}
	if salario > 150000.00 {
		porcentaje += 0.10
	}
	if porcentaje == 0.00 {
		return 0.00
	}
	return salario * porcentaje
}
