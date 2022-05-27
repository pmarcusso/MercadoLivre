package exercicios

import "fmt"

/**
Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
qual possamos passar N quantidade de números inteiros e devolva a média.
Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
*/
func Exercicio2(notas ...float64) (media float64) {

	somaDosValores := 0.0

	for _, valor := range notas {
		somaDosValores += valor
	}

	media = somaDosValores / float64(len(notas))
	fmt.Printf("Exercício 2: O total da média de %d alunos é de %.2f\n", len(notas), media)
	return media
}
