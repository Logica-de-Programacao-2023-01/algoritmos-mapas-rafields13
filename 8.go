// Escreva uma função que receba um mapa onde as chaves são nomes de pessoas e os
// valores são as quantias de dinheiro que cada pessoa gastou em uma conta
// compartilhada. A função deve calcular o valor que cada pessoa deve receber ou
// pagar para igualar as despesas.

package main

import "fmt"

func equalizeExpenses(expenses map[string]float64) map[string]float64 {
	total := 0.0
	numPeople := len(expenses)
	equalAmount := total / float64(numPeople)
	result := make(map[string]float64)

	for _, expense := range expenses {
		total += expense
	}

	for person, expense := range expenses {
		amount := expense - equalAmount
		result[person] = amount
	}

	return result
}

func main() {
	expenses := map[string]float64{
		"Rafael": 10.0,
		"Júlia":  20.0,
		"Caique": 15.0,
	}

	equalizedExpenses := equalizeExpenses(expenses)

	fmt.Println("Equalized Expenses:")
	for person, amount := range equalizedExpenses {
		fmt.Printf("%s: %.2f\n", person, amount)
	}
}
