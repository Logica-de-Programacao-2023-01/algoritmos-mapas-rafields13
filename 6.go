// Escreva uma função que receba uma lista de mapas, onde cada mapa contém a
// contagem de palavras de um texto, e retorne um único mapa contendo a soma de
// todas as contagens.

package main

import "fmt"

func sumCounts(wordsCounts []map[string]int) map[string]int {
	counts := make(map[string]int)

	for _, count := range wordsCounts {
		for word, frequency := range count {
			counts[word] += frequency
		}
	}

	return counts
}

func main() {
	wordCounts := []map[string]int{
		{"Hello,": 1, "world!": 1},
		{"Olá,": 1, "mundo!": 1},
		{"¡Hola,": 1, "mundo!": 1},
	}

	counts := sumCounts(wordCounts)
	fmt.Print(counts)
}
