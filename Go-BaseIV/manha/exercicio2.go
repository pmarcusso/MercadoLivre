package manha

import (
	"errors"
	"fmt"
)

func Exercicio2(salario float64) (string, error) {

	if salario < 15000 {
		return "", errors.New("exercicio2 : erro: O salário digitado não está dentro do valor mínimo")
	}
	return fmt.Sprintln("Necessário pagamento do imposto"), nil
}
