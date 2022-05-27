package manha

import (
	"fmt"
)

/**
Repita o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de
erro receba como parâmetro o valor de "salario", indicando que não atinge o mínimo
tributável (a mensagem exibida pelo console deve dizer : "erro: o mínimo tributável é 15.000 e
o salário informado é: [salario]”, onde [salario] é o valor do tipo int passado pelo parâmetro).
*
*/

func Exercicio3(salario float64) (string, error) {
	if salario < 15000 {
		return "", fmt.Errorf("exercicio3 : erro: o mínimo tributável é 15.000 e o salário informado é: %.2f", salario)
	}
	return fmt.Sprintln("Necessário pagamento do imposto"), nil
}
