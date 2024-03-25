package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 19 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 16 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	fmt.Println("---—----------")
	// FIM DOS ARITMETICOS

	// OPERADORES RELACIONAIS
	// < > <= >= == !=
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM OPERADORES RELACIONAIS

	// LOGICOS
	// && || !
	fmt.Println("---—----------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	fmt.Println("---—----------")
	// FIM DOS OPERADORES LOGICOS

	// OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)
	numero--
	numero -= 20 // numero = numero - 29
	numero *= 3  // numero = numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	fmt.Println("---—----------")
	// FIM OPERADORES UNARIOS

	// nao existe operador ternario
}
