package exercicios

import (
	"fmt"
)

var palavra = "MercadoLivre"

func Exercicio1() {

	for i, letra := range palavra {
		fmt.Println(i, string(letra))
	}

	fmt.Println("A palavra tem", len(palavra))

}
