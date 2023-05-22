// Escreva uma função que receba uma string contendo uma frase e retorne um mapa
// onde as chaves são as palavras encontradas na frase e os valores são mapas
// contendo a contagem de cada letra em cada palavra.

package main

import (
	"fmt"
	"strings"
)

func characteristicsSentences(text string) map[string]map[string]int {
	words := strings.Fields(text)
	characteristics := make(map[string]map[string]int)

	for _, word := range words {
		letterCounter := make(map[string]int)
		for _, character := range word {
			letterCounter[string(character)]++
		}

		characteristics[word] = letterCounter
	}

	return characteristics
}

func main() {
	text := "Hello, world!"

	characteristics := characteristicsSentences(text)
	fmt.Print(characteristics)
}
