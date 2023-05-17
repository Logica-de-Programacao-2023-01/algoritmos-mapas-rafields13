// Escreva uma função que receba um slice de palavras e retorne um mapa onde as
// chaves são palavras que são anagramas entre si e os valores são os grupos de
// palavras anagramas.

package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagramsOccurrences(words []string) map[string][]string {
	anagramsGroups := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		anagramsGroups[sortedWord] = append(anagramsGroups[sortedWord], word)
	}

	return anagramsGroups
}

func sortString(term string) string {
	sorted := strings.Split(term, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func main() {
	words := []string{"amor", "roma", "alegria", "regalia", "galeira", "alergia",
		"cantiga", "catinga", "carro", "corar", "muro", "rumo",
	}

	anagramsGroups := anagramsOccurrences(words)

	for _, anagrams := range anagramsGroups {
		fmt.Print(anagrams)
	}
}
