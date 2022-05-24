package tarde

import "fmt"

/**
Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
para as funções:
- A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
senha
E devem implementar as seguintes funcionalidades:
- mudar o nome: me permite mudar o nome e sobrenome
- mudar a idade: me permite mudar a idade
- mudar e-mail: me permite mudar o e-mail
- mudar a senha: me permite mudar a senha
*/
type usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	email     string
	senha     string
}

func (u Usuario) mudarNomeESobrenome(nome string, sobrenome string, usuario *Usuario) {
	fmt.Println("Usuario antes mudanca:", "Nome:", usuario.Nome, "Sobrenome:", usuario.Sobrenome)
	usuario.Nome = nome
	usuario.Sobrenome = sobrenome
	fmt.Println("Usuario após mudanca:", "Nome:", usuario.Nome, "Sobrenome:", usuario.Sobrenome)
}

func (u Usuario) mudarIdade(idade int, usuario *Usuario) {
	fmt.Println("Usuario antes mudanca:", "Idade:", usuario.Idade)
	usuario.Idade = idade
	fmt.Println("Usuario após mudanca:", "Idade:", usuario.Idade)
}

func (u Usuario) mudarEmail(email string, usuario *Usuario) {
	fmt.Println("Usuario antes mudanca:", "Email:", usuario.email)
	usuario.email = email
	fmt.Println("Usuario antes mudanca:", "Email:", usuario.email)
}

func (u Usuario) mudarSenha(senha string, usuario *Usuario) {
	fmt.Println("Usuario antes mudanca:", "Senha:", usuario.senha)
	usuario.senha = senha
	fmt.Println("Usuario antes mudanca:", "Senha:", usuario.senha)
}

func Exercicio1() {

	usuario := Usuario{Nome: "Paulo", Sobrenome: "Marcusso", Idade: 30, email: "123@meli.com", senha: "123"}
	usuario.mudarNomeESobrenome("Angelo", "Angelical", &usuario)
	usuario.mudarIdade(31, &usuario)
	usuario.mudarEmail("angelo@meli.com", &usuario)
	usuario.mudarSenha("456", &usuario)
	fmt.Print(usuario)
}
