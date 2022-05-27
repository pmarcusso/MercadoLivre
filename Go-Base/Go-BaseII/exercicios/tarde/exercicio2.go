package tarde

import (
	"log"
)

type loja struct {
	Nome     string
	Produtos []produto
}
type produto struct {
	TipoProduto string
	Nome        string
	Preco       float64
	Quantidade  int
}

type Produto interface {
	CalcularCusto()
}

type Ecommerce interface {
	Total()
	Adicionar(produto produto)
}

func Exercicio2() {
	novaLoja := novaLoja("Ebazar")
	novoProduto := novoProduto("pequeno", "prato", 20.0, 5)
	listaDeProdutos := novaLoja.Adicionar(novoProduto)

	log.Printf("Loja %v - Produtos: %v", novaLoja, listaDeProdutos)
}

func novoProduto(tipoProduto string, nome string, preco float64, quantidade int) produto {
	novoProduto := produto{tipoProduto, nome, preco, quantidade}
	log.Printf("Criando produto: %v", novoProduto)

	return novoProduto
}

func novaLoja(nome string) loja {

	loja := loja{Nome: nome}
	log.Printf("Criando loja %v", loja)
	return loja
}

func (p produto) CalcularCusto() {

	if p.TipoProduto == "pequeno" {

	} else if p.TipoProduto == "medio" {

	} else if p.TipoProduto == "grande" {

	}
}

func (l *loja) Adicionar(produto produto) []produto {

	produtoAdicionado := append(l.Produtos, produto)
	log.Printf("Produto adicionado: %v na loja %v", produtoAdicionado, l)
	return produtoAdicionado
}
