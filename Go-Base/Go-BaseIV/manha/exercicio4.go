package manha

import (
	"errors"
	"fmt"
)

/**
Vamos fazer com que nosso programa seja um pouco mais complexo e útil.
1. Desenvolva as funções necessárias para permitir que a empresa calcule:
a) Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
- A função receberá as horas trabalhadas no mês e o valor da hora como
parâmetro.
- Esta função deve retornar mais de um valor (salário calculado e erro).
- No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser
descontados como imposto.

2

- Se o número de horas mensais inseridas for menor que 80 ou um número
negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador
não pode ter trabalhado menos de 80 horas por mês".

2. Desenvolva o código necessário para cumprir as funcionalidades solicitadas, usando
“errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de realizar as validações dos
retornos de erro em sua função “main()”.
*
*/

func Exercicio4() {
	salario, err := calcSalario(20, 100)
	if err != nil {
		fmt.Println("Operacão inválida:", err)
	}

	fmt.Println("O salário atual é de:", salario)
}

func calcSalario(horasTrabalhadas float64, valorHora float64) (float64, error) {

	var imposto float64
	var salarioComImposto float64

	salarioSemImposto := horasTrabalhadas * valorHora

	if horasTrabalhadas < 80.0 {
		err := errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês ")
		return salarioSemImposto, err
	}

	if salarioSemImposto >= 20000 {
		imposto = salarioSemImposto * 0.10
		salarioComImposto = salarioSemImposto - imposto
		fmt.Printf("O salário de %.2f terá um imposto de %.2f\n", salarioSemImposto, imposto)
	} else {
		imposto = 0.0
		salarioComImposto = salarioSemImposto - imposto
	}

	return salarioComImposto, nil
}
