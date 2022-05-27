package tarde

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
O mesmo escritório do exercício anterior, solicita uma funcionalidade para poder
cadastrar dados de novos clientes. Os dados necessários para cadastrar um cliente são:
- Arquivo
- Nome e sobrenome
- RG
- Número de telefone
- Endereço

● Tarefa 1: O número do arquivo deve ser atribuído ou gerado separadamente e antes
da cobrança das despesas restantes. Desenvolva e implemente uma função para
gerar um ID que você usará posteriormente para atribuí-lo como um valor a “Arquivo”.
Se por algum motivo esta função retornar o valor "nil", deve gerar um panic que
interrompa a execução e aborte

2
● Tarefa 2: Antes de cadastrar um cliente, você deve verificar se o cliente já existe. Para
fazer isso, você precisa ler os dados de um arquivo .txt. Em algum lugar do seu
código, implemente a função para ler um arquivo chamado “customers.txt” (como no
exercício anterior, este arquivo não existe, então qualquer função que tente lê-lo
retornará um erro). Você deve lidar adequadamente com esse erro, como vimos até
agora. Esse erro deve:

1. Gerar um panic;
2. Imprimir no console a mensagem: “erro: o arquivo indicado não foi encontrado ou
está danificado”, e continuar com a execução do programa normalmente.

Requisitos gerais:
● Use recover para recuperar o valor dos panics que podem surgir (exceto na tarefa 1).
● Lembre-se de realizar as validações necessárias para cada retorno que possa conter
um valor de erro (por exemplo, aqueles que tentam ler arquivos).
Crie um erro, personalizando-o ao seu gosto, utilizando qualquer uma das funções
que o GO disponibiliza para ele (ele também deve realizar a validação relevante para
o caso de erro retornado).
*/

type Cliente struct {
	Arquivo   string
	Nome      string
	Sobrenome string
	RG        string
	Telefone  string
	Endereco  string
}

func Exercicio2() {
	cliente := Cliente{Arquivo: gerarID()}
	arquivo := gerarArquivo("customers.txt")
	_, err := adicionarCliente(arquivo.Name(), &cliente)

	if err != nil {
		panic("Cliente já cadastrado")
	}

}

func checkID(arquivo *os.File, idCliente string) (bool, string) {

	defer func() {
		fmt.Println("Falha ao checar o id!")
	}()

	arquivoLido, err := ioutil.ReadFile(arquivo.Name())

	if err != nil {
		panic(err)
	}

	palavra := string(arquivoLido)

	if strings.Contains(palavra, idCliente) {
		return true, idCliente
	}

	log.Println("ID não existe")
	return false, idCliente
}

func adicionarCliente(nomeArquivo string, c *Cliente) (*os.File, error) {

	arquivo, err := os.OpenFile(nomeArquivo, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	defer arquivo.Close()

	if err != nil {
		panic("erro: o arquivo indicado não foi encontrado ou está danificado")
	}

	idExiste, _ := checkID(arquivo, c.Arquivo)
	if idExiste {
		log.Printf("O cliente com o id %s já está cadastrado", idExiste)
		return arquivo, nil
	}

	_, err = fmt.Fprintf(arquivo, c.Arquivo+"\n")

	if err != nil {
		fmt.Println(err)
		arquivo.Close()
		return nil, err
	}

	defer arquivo.Close()

	fmt.Println("Novo id inserido no arquivo", nomeArquivo)

	return arquivo, nil
}

func gerarID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func gerarArquivo(nomeArquivo string) *os.File {

	//checa se o arquivo já existe
	arquivo, err := os.Open(nomeArquivo)

	if err != nil {

		novoArquivo, err := os.Create(nomeArquivo)
		defer novoArquivo.Close()

		if err != nil {
			panic(err)
		}

		return novoArquivo

	}

	defer arquivo.Close()
	return arquivo
}
