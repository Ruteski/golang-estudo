package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Jhon", "Doe", 35, 179}
	fmt.Println(p1)

	e1 := estudante{p1, "direito", "ufpr"}
	fmt.Println(e1)
	fmt.Println(e1.nome)

	// dessa forma nao funciona, sempre tem que criar o pessoa antes de criar o estudante
	// e2 := estudante{"Mary", "Doe", 30, 165, "sistemas", "ufpr"}
}
