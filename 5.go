// Escreva uma função que receba uma string e retorne um mapa onde as chaves são
// os caracteres presentes na string e os valores são a frequência de cada
// caractere.

package main

import "fmt"

func charactersCounter(text string) map[string]int {
	occurrences := make(map[string]int)

	for _, character := range text {
		occurrences[string(character)]++
	}

	return occurrences
}

func main() {
	text := "Hello, world!"

	occurrences := charactersCounter(text)
	fmt.Print(occurrences)
}
