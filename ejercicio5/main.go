package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		perro = "perro"
		gato  = "gato"
	)
	animalPerro, err := Animal(perro)
	animalGato, err := Animal(gato)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	var cantidad int
	cantidad += animalPerro(5)
	cantidad += animalGato(8)

	fmt.Println("hola: ", cantidad)
}

func Animal(animalType string) (func(int) int, error) {
	switch animalType {
	case "perro":
		return perroCal, nil
	}

	return nil, errors.New("error")
}

func perroCal(cantidad int) int {
	kgA := 10
	return cantidad * kgA
}
