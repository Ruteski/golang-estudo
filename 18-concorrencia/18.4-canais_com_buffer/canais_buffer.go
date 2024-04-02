package main

import (
	"fmt"
)

func main() {
	// buffer com capacidade 2
	canal := make(chan string, 2)
	canal <- "Ola mundo!"
	canal <- "Programando com Go"
	//canal <- "Terceiro valor" // nesse caso ocorre o deadlock, pq o buffer é de capacidade 2

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
