// Escreva uma função que receba um mapa com valores inteiros e retorne a soma de
// todos os valores.

package main

import "fmt"

func sumValues(values map[int]int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func main() {
	values := map[int]int{
		1: 2,
		2: 2,
		3: 1,
		4: 1,
		5: 1,
	}
	sum := sumValues(values)
	fmt.Print(sum)
}
