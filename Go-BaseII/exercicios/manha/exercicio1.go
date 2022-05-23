package exercicios

import "fmt"

/**Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
imposto de um salário.
Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
caso será descontado mais 10%.
*/
func Exercicio1(salario float64) (descontoImposto float64) {
	if salario < 50000.00 {
		descontoImposto = 0.0
	} else if salario >= 50000.00 && salario <= 150000.00 {
		descontoImposto = salario * 0.17
	} else {
		descontoImposto = salario * 0.27
	}
	fmt.Printf("Exercício 1: O desconto do salário é de %.2f\n ", descontoImposto)
	return descontoImposto
}
