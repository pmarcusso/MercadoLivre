package manha

import "fmt"

/**
1. Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
tipo “int”.
2. Crie um erro personalizado com uma struct que implemente “Error()” com a
mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
console a mensagem “Necessário pagamento de imposto”.
*/

type erro struct{}

func (e *erro) Error() string {
	return "Exercicio1 : erro: O salário digitado não está dentro do valor mínimo"
}

func Exercicio1(salario float64) (string, error) {
	if salario < 15000 {
		return "", &erro{}
	}
	return fmt.Sprintln("Necessário pagamento do imposto"), nil
}
