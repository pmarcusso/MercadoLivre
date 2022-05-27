package exercicios

import "fmt"

/**
Faça uma aplicação que contenha uma variável com o número do mês.
1. Dependendo do número, imprima o mês correspondente em texto.
2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você
escolheria e porquê?
Ex: 7 de julho.
*/

func Exercicio3() {

	numero := 12

	var meses = map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Marco",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	for i, mes := range meses {

		if numero == i {
			fmt.Println(mes)
		}
	}
}
