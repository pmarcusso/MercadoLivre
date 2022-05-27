package tarde

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Uma empresa de sistemas precisa analisar que algoritmos de ordenamento utilizar para seus
serviços.
Para eles é necessário instanciar 3 arrays com valores aleatórios não ordenados
- uma matriz de inteiros com 100 valores
- uma matriz de inteiros com 1000 valores
- uma matriz de inteiros com 10.000 valores

Para instanciar as variáveis, utilize o rand:
variavel := rand.Perm(100)
variave2 := rand.Perm(1000)
variave3 := rand.Perm(10000)

Cada um deve ser ordenado por:

- Inserção
- grupo
- seleção

Uma rotina para cada execução de classificação
Tenho que esperar terminar a ordenação de 100 números para continuar com a ordenação de
1000 e depois a ordenação de 10000.
Por fim, devo medir o tempo de cada um e mostrar o resultado na tela, para saber qual
ordenação ficou melhor para cada arranjo.
*/

func Exercicio4() {

	variavel := rand.Perm(100)
	variave2 := rand.Perm(1000)
	variave3 := rand.Perm(10000)

	var functions []func([]int, chan<- string)
	functions = append(functions, SelectionSort, InsertionSort)

	sortList(variavel, functions)
	sortList(variave2, functions)
	sortList(variave3, functions)
}

func InsertionSort(array []int, c chan<- string) {
	start := time.Now()

	var tamanho = len(array)
	for i := 1; i < tamanho; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}

	timeExecution := time.Since(start).Seconds()
	result := fmt.Sprintf("Insertion Sort: Quantidade de números %d inteiros: %.6fs",
		tamanho, timeExecution)
	c <- result
}

func SelectionSort(array []int, c chan<- string) {
	start := time.Now()

	var tamanho = len(array)
	for i := 0; i < tamanho; i++ {
		var minIdx = i
		for j := i; j < tamanho; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[i], array[minIdx] = array[minIdx], array[i]
	}

	timeExecution := time.Since(start).Seconds()
	result := fmt.Sprintf("Selection Sort: Quantidade de números %d inteiros: %.6fs",
		tamanho, timeExecution)
	c <- result
}

func sortList(array []int, funcoes []func([]int, chan<- string)) {
	c := make(chan string)

	for _, funcao := range funcoes {
		go funcao(array, c)
	}

	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}
	fmt.Println()
}
