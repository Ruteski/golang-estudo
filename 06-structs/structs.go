package main

import "fmt"

type usuario struct {
	nome     string
	email    string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint16
}

func main() {
	fmt.Println("Arquivo structs")

	var user usuario
	fmt.Println(user)

	user.email = "sdf@asdf.com"
	user.nome = "Jhon Doe"
	user.idade = 35
	fmt.Println(user)

	enderecoExemplo := endereco{"rua dos alfineiros", 35}

	user2 := usuario{"Mary Doe", "marydoe@asdf.com", 32, enderecoExemplo}
	fmt.Println(user2)

	user3 := usuario{nome: "Jane Doe"}
	fmt.Println(user3)
}
