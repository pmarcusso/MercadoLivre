package tarde

import (
	"fmt"
)

/**
Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total
dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade,
eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

Precisamos de 3 estruturas:
- Produtos: nome, preço, quantidade.
- Serviços: nome, preço, minutos trabalhados.
- Manutenção: nome, preço.
Precisamos de 3 funções:
- Somar Produtos: recebe um array de produto e devolve o preço total (preço *
quantidade).
- Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
tivesse trabalhado meia hora).
- Somar Manutenção: recebe um array de manutenção e devolve o preço total.

Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na
tela (somando o total dos 3).
*/

type produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}
type servico struct {
	Nome    string
	Preco   float64
	Minutos int
}
type manutencao struct {
	Nome  string
	Preco float64
}

func custoManutencao(ch chan float64, manutencoes ...manutencao) {
	var custoTotal float64

	for _, value := range manutencoes {
		custoTotal += value.Preco
	}

	fmt.Println("Custo total manutencão:", custoTotal)

	ch <- custoTotal

	//close(ch)
}

func custoProdutos(ch chan float64, produtos ...produto) {
	var custoTotal float64

	for _, value := range produtos {
		custo := value.Preco * float64(value.Quantidade)
		custoTotal += custo
	}

	fmt.Println("Custo total produtos:", custoTotal)

	ch <- custoTotal

	//close(ch)
}

func custoServicos(ch chan float64, servico ...servico) {
	var custoTotal float64

	for _, value := range servico {
		var mediaHorasTrabalhadas = float64(value.Minutos) / 60
		if mediaHorasTrabalhadas < 0.5 {
			mediaHorasTrabalhadas = 0.5
		}
		custo := mediaHorasTrabalhadas * value.Preco
		custoTotal += custo
	}

	fmt.Println("Custo total servicos:", custoTotal)

	ch <- custoTotal

	//close(ch)
}

func Exercicio3() {
	novoProduto1 := produto{Nome: "Caneca", Preco: 10.0, Quantidade: 3}
	novoProduto2 := produto{Nome: "Caf", Preco: 15.0, Quantidade: 3}
	novoServico1 := servico{Nome: "Troca de hardware", Preco: 50.0, Minutos: 60}
	novoServico2 := servico{Nome: "Manutencao de isqueiros", Preco: 100.0, Minutos: 25}
	novaManutencao1 := manutencao{Nome: "Manutencao cadeira", Preco: 45.0}
	novaManutencao2 := manutencao{Nome: "Manutencao mesa", Preco: 45.0}

	produtos := make([]produto, 0)
	manutencao := make([]manutencao, 0)
	servicos := make([]servico, 0)

	produtos = append(produtos, novoProduto1, novoProduto2)
	servicos = append(servicos, novoServico1, novoServico2)
	manutencao = append(manutencao, novaManutencao1, novaManutencao2)

	ch := make(chan float64)

	go custoManutencao(ch, manutencao...)
	go custoProdutos(ch, produtos...)
	go custoServicos(ch, servicos...)

	var total float64 = 0

	for i := 0; i < 3; i++ {
		total += <-ch
	}
	//Inicia os calculos
	fmt.Printf("O total do serviço completo é de R$%.2f\n", total)

}
