package main

import "fmt"

// nao precisa especificar a quantidade de parametros que a func vai receber

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	total := soma(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	escrever("numero:", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
