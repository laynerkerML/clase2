package main

import "fmt"

type Matrix struct {
	Alto  float64
	Ancho float64
}

func main() {
	fmt.Println("hola")
}

func (m *Matrix) setMatrix(parametos Matrix) {
	m.Alto = parametos.Alto
	m.Ancho = parametos.Ancho
}
