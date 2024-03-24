package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

/*
	go mod init <nome do modulo>

	go get <pacote> // github.com/badoux/checkmail

	go mod tidy // remove todas as importacoes nao usadas
*/

func main() {
	fmt.Println("Escrevendo do arquivo main.")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("123@")
	fmt.Println(erro)
}
