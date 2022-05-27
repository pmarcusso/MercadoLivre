package exercicios

import (
	"errors"
	"fmt"
)

/**
Um abrigo de animais precisa descobrir quanta comida comprar para os animais de
estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão
é que haja muito mais animais para abrigar.
1. Cães precisam de 10 kg de alimento
2. Gatos 5 kg
3. Hamster 250 gramas.
4. Tarântula 150 gramas.

Precisamos:
1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto
com o animal especificado e que retorne uma função com uma mensagem (caso não
exista o animal)
2. Uma função para cada animal que calcule a quantidade de alimento com base na
quantidade necessária do animal digitado.
*/

const (
	cachorro  = "cachorro"
	gato      = "gato"
	tarantula = "tarantula"
	hamster   = "hamster"
)

func Animal(especie string) (func(quantidade int) int, error) {

	switch especie {
	case "tarantula":
		return calcTarantula, nil
	case "hamster":
		return calcHamster, nil
	case "cachorro":
		return calcCachorro, nil
	case "gato":
		return calcGato, nil
	default:
		return nil, errors.New("Espécie inválida")
	}
}

func calcTarantula(quantidade int) int {
	return quantidade * 150
}

func calcCachorro(quantidade int) int {
	return quantidade * 10000
}

func calcGato(quantidade int) int {
	return quantidade * 5000
}

func calcHamster(quantidade int) int {
	return quantidade * 5000
}

func Exercicio5() {
	dfunc, _ := Animal(cachorro)
	fmt.Printf("quandidade de alimento do(s) cachorro(s) (em gramas):%d gramas\n", dfunc(5))
	cfunc, _ := Animal(gato)
	fmt.Printf("quandidade de alimento do(s) gato(s) (em gramas): %d gramas\n", cfunc(8))
	_, err := Animal("invalid animal")
	if err != nil {
		fmt.Println("animal inválido")
	}
}
