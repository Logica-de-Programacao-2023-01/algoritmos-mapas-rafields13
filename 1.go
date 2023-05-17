// Escreva uma função que conte a ocorrência de cada palavra em uma string e
// retorne um mapa onde as chaves são as palavras encontradas e os valores são o
// número de ocorrências de cada palavra.

package main

import (
	"fmt"
	"strings"
)

func wordsCounter(text string) map[string]int {
	words := strings.Fields(text)
	occurrences := make(map[string]int)

	for _, word := range words {
		occurrences[word]++
	}

	return occurrences
}

func main() {
	text := "Hello, world!"

	occurrences := wordsCounter(text)
	fmt.Print(occurrences)
}
