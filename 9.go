// Escreva uma função que gere a sequência de Fibonacci até um determinado número
// n e retorne um mapa onde as chaves são os índices da sequência e os valores
// são os números correspondentes.

package main

import "fmt"

func fibonacciSequence(n int) map[int]int {
	sequence := make(map[int]int)

	sequence[0] = 0
	if n > 0 {
		sequence[1] = 1
	}

	for i := 2; i <= n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence
}

func main() {
	n := 10
	fibSeq := fibonacciSequence(n)

	fmt.Println("Fibonacci Sequence:")
	for i, num := range fibSeq {
		fmt.Printf("Index %d: %d\n", i, num)
	}
}
