// Escreva uma função que receba um slice de inteiros e retorne um mapa onde as
// chaves são os pares de números encontrados no slice e os valores são as
// frequências de cada par.

package main

import "fmt"

func findPairFrequencies(numbers []int) map[[2]int]int {
	pairFrequencies := make(map[[2]int]int)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			pairFrequencies[pair]++
		}
	}

	return pairFrequencies
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 1, 2, 3, 4}
	pairFreq := findPairFrequencies(numbers)

	fmt.Println("Pair Frequencies:")
	for pair, freq := range pairFreq {
		fmt.Printf("%v: %d\n", pair, freq)
	}
}
