package main

import "fmt"

// funcao init - 1 por arquivo
var n int

func init() {
	fmt.Println("Funcao init sendo executada")
	n = 10
}

func main() {
	fmt.Println("Funcao main sendo executada")
	fmt.Println(n)
}
