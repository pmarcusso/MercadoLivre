package tarde

import (
	"errors"
	"fmt"
)

/**
Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
mesmo endereço de memória no main do programa e nas funções.

Estruturas necessárias:
- Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
- Produto: Nome, preço, quantidade.
Algumas funções necessárias:
- Novo produto: recebe nome e preço, e retorna um produto.
- Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona
o produto ao usuário.
- Deletar produtos: recebe um usuário, apaga os produtos do usuário.
*/

type Usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	email     string
	senha     string
	Produto   []Produto
}

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func novoProduto(nome string, preco float64, quantidade int) Produto {

	novoProduto := Produto{Nome: nome, Preco: preco, Quantidade: quantidade}

	fmt.Println("Novo produto adicionado: ", novoProduto)

	return novoProduto
}

func adicionarProduto(usuario *Usuario, produtos ...Produto) {
	if usuario == nil || len(produtos) == 0 {
		errors.New("É preciso fazer o input do usuario e/ou dos produtos")
	}

	for _, valor := range produtos {
		usuario.Produto = append(usuario.Produto, valor)
	}

	fmt.Printf("Usuario %+v adicionou os produtos %+v\n", usuario.Nome, produtos)
}

func removerProdutos(usuario *Usuario) {
	usuario.Produto = nil
}

func Exercicio2() {
	usuario1 := Usuario{"Paulo Henrique", "Marcusso", 30, "meli@meli", "senha", nil}
	fmt.Println("Usuario", usuario1)
	novoProduto1 := novoProduto("Macbook Pro", 25000.0, 15)
	novoProduto2 := novoProduto("Jumanji", 700.0, 32)
	adicionarProduto(&usuario1, novoProduto1, novoProduto2)
	fmt.Println(usuario1)

	removerProdutos(&usuario1)
	fmt.Println("Removendo produtos do usuario", usuario1)
}
