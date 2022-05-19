package exercicios

import "fmt"

var (
	temperatura float64 = 5.5
	umidade     float64 = 60.7
	pressaoAtm  float64 = 1
)

func Exercicio2() {
	fmt.Println("Exercício 2: ", "A temperatura está em", temperatura, "graus Celsius", "com umidade de ", umidade, "%e com pressão atmosférica de", pressaoAtm, "atm")

}
