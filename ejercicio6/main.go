package main

import (
	"errors"
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func main() {
	searchDNI := 7667676
	resultSearch, err := detalles(searchDNI)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Alumno: ", resultSearch)
}

func detalles(DNI int) (Alumno, error) {
	result := Alumno{}
	alumnos := make([]Alumno, 2)
	alumnos[0] = Alumno{Nombre: "Laynerker", Apellido: "Guerrero", DNI: 95950356, Fecha: "13-12-2021"}
	alumnos[1] = Alumno{Nombre: "Yissel", Apellido: "Garcia", DNI: 95959163, Fecha: "14-12-2021"}
	for _, alumn := range alumnos {
		if alumn.DNI == DNI {
			result = alumn
		}
	}
	if result == (Alumno{}) {
		return result, errors.New("No existe el alumno")
	}

	return result, nil
}
