package tarde

import (
	"encoding/json"
	"fmt"
	"time"
)

/**
Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os
detalhes dos dados de cada um deles, conforme o exemplo abaixo:
Nome: [Nome do aluno]
Sobrenome: [Sobrenome do aluno]
RG: [RG do aluno]
Data de admissão: [Data de admissão do aluno]

Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos
alunos.
Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG,
Data e que tenha um método de detalhamento
*/

//   OBS : GO trabalha com uma forma diferente de data então é preciso dar uma olhada na documentação

type Aluno struct {
	Nome         string    `json:"nome"`
	Sobrenome    string    `json:"sobrenome"`
	RG           string    `json:"rg"`
	DataAdmissao time.Time `json:"data_admissao"`
}

func (a Aluno) Detalhe() {
	fmt.Printf("Aluno:\t Nome: %s\t Sobrenome:%s\t RG: %s\t DataAdmissao: %s\t \n", a.Nome, a.Sobrenome, a.RG, a.DataAdmissao)
}

func Exercicio1() {
	aluno1 := Aluno{"Paulo Henrique", "Marcusso", "11111111", time.Date(1980, 01, 15, 0, 0, 0, 0, time.UTC)}

	//Transformando em JSON
	alunoJson, err := json.Marshal(aluno1)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(alunoJson))
	aluno1.Detalhe()
}
