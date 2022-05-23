package exercicios

import (
	"errors"
	"fmt"
)

/**Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
de suas notas.
Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
indicado na função anterior
*/
const (
	minimum = "minimun"
	average = "average"
	maximum = "maximum"
)

func calcMin(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("Nenhum valor inputado!")
	}

	min := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}

	return min, nil
}

func calcMax(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("Nenhum valor inputado!")
	}

	max := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}

	return max, nil
}

func calcAvg(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0.0, errors.New("Nenhum valor inputado!")
	}

	avr := 0.0

	for _, value := range values {
		avr += value
	}

	return avr / float64(len(values)), nil
}

func getFunc(t string) (func(values ...float64) (float64, error), error) {
	if t == minimum {
		return calcMin, nil
	}

	if t == average {
		return calcAvg, nil
	}

	if t == maximum {
		return calcMax, nil
	}

	return nil, errors.New("invalid function type")
}

func Exercicio4() {

	avgFunc, _ := getFunc(average)
	maxFunc, _ := getFunc(maximum)
	minFunc, _ := getFunc(minimum)

	min, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("max: %.2f\n", max)
}
