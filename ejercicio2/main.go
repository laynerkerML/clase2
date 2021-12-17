package main

import (
	"errors"
	"fmt"
)

func main() {
	calulo, err := calculaPromedios(1, 2, 3, 4, 5, 6, -7, 7, 8, 8, 8)
	if err == nil {
		fmt.Println("calculo: ", calulo)
	} else {
		fmt.Println("error: ", err)
	}

}

func calculaPromedios(notas ...int) (float64, error) {
	sumeNotas := 0
	for _, n := range notas {
		if n <= -1 {
			return 0.00, errors.New("error hay un valor negativo")
		}
		sumeNotas += n
	}
	return float64(sumeNotas) / float64(len(notas)), nil
}
