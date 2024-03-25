package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel3 int
	var ponteiro *int // o valor zerado de um ponteiro é nil
	fmt.Println(variavel3, ponteiro)

	variavel3 = 20
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // o '*' serve pra mostrar o valor que esta dentro do ponteiro, desreferenciacao

	variavel3 = 35
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}
