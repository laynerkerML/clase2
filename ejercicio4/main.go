package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		minimo   = "minimo"
		promedio = "promedio"
		maximo   = "maximo"
	)
	minFunc := operacion(minimo)
	promFunc := operacion(promedio)
	maxFunc := operacion(maximo)

	valorMinimo, err := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio, err := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo, err := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(minimo, ": ", valorMinimo)
	fmt.Println(promedio, ": ", valorPromedio)
	fmt.Println(maximo, ": ", valorMaximo)
}

func operacion(calType string) func(notas ...int) (float64, error) {
	switch calType {
	case "minimo":
		return minimoCal
	case "maximo":
		return maximoCal
	case "promedio":
		return promerioCal
	}
	return nil
}

func minimoCal(notas ...int) (float64, error) {
	numMenor := notas[0]
	for _, nota := range notas {
		if nota <= -1 {
			return 0, errors.New("error")
		}
		if nota < numMenor {
			numMenor = nota
		}
	}
	return float64(numMenor), nil
}

func maximoCal(notas ...int) (float64, error) {
	numMayor := notas[0]
	for _, nota := range notas {
		if nota <= -1 {
			return 0, errors.New("error")
		}
		if nota > numMayor {
			numMayor = nota
		}
	}
	return float64(numMayor), nil
}

func promerioCal(notas ...int) (float64, error) {
	sumaNotas := 0
	for _, nota := range notas {
		if nota <= -1 {
			return 0, errors.New("error")
		}
		sumaNotas += nota
	}
	return float64(sumaNotas) / float64(len(notas)), nil
}
