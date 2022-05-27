package exercicios

import (
	"errors"
)

/**
Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
de horas trabalhadas por mês e na categoria do funcionário.
Se a categoria for C, seu salário é de R$1.000 por hora
Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
mensais
Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
mensais

Calcule o salário dos funcionários conforme as informações abaixo:
- Funcionário de categoria C: 162h mensais
- Funcionário de categoria B: 176h mensais
- Funcionário de categoria A: 172h mensais
*/
func Exercicio3(categoria string, horasTrabalhadas float64) (float64, error) {

	var salario float64

	if categoria == "C" {
		salario = horasTrabalhadas * 1000
		return salario, nil
	} else if categoria == "B" {
		salario = horasTrabalhadas * 1500

		if horasTrabalhadas > 160 {
			salario = salario * (1.2)
		}
		return salario, nil
	} else if categoria == "A" {
		salario = horasTrabalhadas * 3000

		if horasTrabalhadas > 160 {
			salario = salario * (1.5)
		}
		return salario, nil
	}

	return 0.0, errors.New("Categoria Inválida")
}
