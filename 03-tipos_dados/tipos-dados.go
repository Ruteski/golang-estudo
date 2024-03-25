package main

import (
	"errors"
	"fmt"
)

func main() {
	//int, int8, int16, int32, int64
	//uint, uint8, uint16, uint32, uint64 - só aceita numeros positivos

	var numero1 int64 = 1000000000000000
	fmt.Println(numero1)

	numero2 := 1000000000000000
	fmt.Println(numero2)

	// var numero3 uint32 = -10000
	// fmt.Println(numero3)

	// alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12345000000.67
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// string
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	// "char"
	char := 'B' //mostra o numero da tabela ascii correspondente ao character
	fmt.Println(char)

	// =====
	var texto string
	fmt.Println(texto)

	var num int16
	fmt.Println(num)

	// booelan
	var booleano bool = true
	fmt.Println(booleano)

	// tipo erro
	var erro error = errors.New("Erro interno") // como instanciar um erro
	fmt.Println(erro)

}
