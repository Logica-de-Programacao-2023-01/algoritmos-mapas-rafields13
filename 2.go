// Escreva uma função que receba dois mapas e retorne um novo mapa contendo todos
// os elementos dos mapas de entrada. Em caso de chaves duplicadas, o valor do
// segundo mapa deve prevalecer.

package main

import "fmt"

func returnMap(firstMap map[int]int, secondMap map[int]int) map[int]int {
	newMap := make(map[int]int)

	for key, value := range firstMap {
		newMap[key] = value
	}
	for key, value := range secondMap {
		newMap[key] = value
	}
	return newMap
}

func main() {
	firstMap := map[int]int{
		1: 2,
		2: 2,
		3: 1,
		4: 1,
		5: 1,
	}
	secondMap := map[int]int{
		1: 1,
		2: 1,
		3: 1,
		4: 2,
		5: 2,
	}
	newMap := returnMap(firstMap, secondMap)
	fmt.Print(newMap)
}
