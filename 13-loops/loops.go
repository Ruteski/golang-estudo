package main

import (
	"fmt"
)

// nao se usa range em struct

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i:", i)
		//time.Sleep(time.Second)
	}

	fmt.Println("----------------")

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j:", j)
		//time.Sleep(time.Second)
	}
	fmt.Println("----------------")
	nomes := [3]string{"Joao", "Davi", "Lucas"} // array
	nomes2 := []string{"Joao", "Davi", "Lucas"} // slice

	for indice, nome := range nomes {
		fmt.Println("indice:", indice, "nome:", nome)
	}

	fmt.Println("----------------")
	for _, nome := range nomes2 {
		fmt.Println("nome:", nome)
	}

	fmt.Println("----------------")
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, " - código ASCII:", letra, " - valor:", string(letra))
	}

	fmt.Println("----------------")
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
