package exercicios

import (
	"fmt"
)

var palavra = "MercadoLivre"

func Exercicio1() {

	var numero = []int{}
	for i, letra := range palavra {
		fmt.Println(i, string(letra))
		numero = append(numero, i)
	}

	fmt.Println("A palavra tem", len(numero))
}
