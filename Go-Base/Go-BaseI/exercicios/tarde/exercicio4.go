package exercicios

import "fmt"

/**
- Saiba quantos de seus funcionários têm mais de 21 anos.
- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Excluir Pedro do mapa.
*/

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func Exercicio4() {

	qntMaiorQue21Anos := make([]int, 0)

	for key, value := range employees {
		if key == "Benjamin" {
			fmt.Println("A idade é:", value)
		}

		if value > 21 {
			qntMaiorQue21Anos = append(qntMaiorQue21Anos, value)
		}
	}

	fmt.Println("Quantidade de funcionarios com mais de 21 anos é de :", len(qntMaiorQue21Anos))

	employees["Frederico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)
}
