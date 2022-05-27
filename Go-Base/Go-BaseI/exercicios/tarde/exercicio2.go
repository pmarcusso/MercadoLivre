package exercicios

import "fmt"

/**
Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.
*/
func Exercicio2() {
	idade := 30
	empregado := true
	tempoDeAtividadeEmAnos := 2

	if idade >= 22 && empregado && tempoDeAtividadeEmAnos >= 1 {
		fmt.Println("O cliente está apto a receber um empréstimo")
	} else {
		fmt.Println("O cliente não atende a algum requisito")
	}

	salarioAnual := 100.000

	if salarioAnual > 100.000 {
		fmt.Println("O cliente está isento de juros")
	} else {
		fmt.Println("O cliente irá pagar juros sobre o empréstimo")
	}

}
