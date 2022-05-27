package manha

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

/**
Exercício 1 - Guardar arquivo
Uma empresa que vende produtos de limpeza precisa de:
1. Implementar uma funcionalidade para guardar um arquivo de texto, com a informação
de produtos comprados, separados por ponto e vírgula (csv).
2. Deve possuir o ID do produto, preço e a quantidade.
3. Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

A mesma empresa precisa ler o arquivo armazenado, para isso exige que:
- Seja impresso na tela os valores tabelados, com título ( à esquerda para o ID e à
direita para o Preço e Quantidade), o preço, a quantidade e abaixo do preço o total
deve ser exibido (somando preço por quantidade)
*/

type Produto struct {
	ID         int
	Preco      float64
	Quantidade int
}

func Exercicio1() {

	novoProduto1 := Produto{1, 10.0, 3}
	novoProduto2 := Produto{2, 20.0, 6}
	novoProduto3 := Produto{3, 30.0, 9}

	escreverArquivo(novoProduto1, novoProduto2, novoProduto3)
}

func escreverArquivo(produtos ...Produto) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	_, err = arquivo.WriteString("ID" + "\t\t\t\t" + "PRECO" + "\t\t\t\t\t" + "QUANTIDADE" + "\t\t\n")

	if err != nil {
		fmt.Println(err)
	}
	var precoTotal float64
	for i := 0; i < len(produtos); i++ {

		registrarProduto(produtos[i])
		quantidadeConvertida := float64(produtos[i].Quantidade)
		precoProduto := produtos[i].Preco
		precoTotal += precoProduto * quantidadeConvertida

	}

	precoTotalConvertido := fmt.Sprintf("%f", precoTotal)
	_, err = arquivo.WriteString("\t\t\t\t" + precoTotalConvertido + "\n")

	if err != nil {
		fmt.Println(err)
	}
	arquivo.Close()
	imprimirLogs()
}

func registrarProduto(produto Produto) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	idConvertido := strconv.Itoa(produto.ID)
	precoConvertido := fmt.Sprintf("%f", produto.Preco)
	quantidadeConvertida := strconv.Itoa(produto.Quantidade)

	_, err = arquivo.WriteString(idConvertido + "\t\t\t\t" + precoConvertido + "\t\t" + "\t\t" + quantidadeConvertida + ";\n")
	if err != nil {
		fmt.Println(err)
	}
}

func imprimirLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
