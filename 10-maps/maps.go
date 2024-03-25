package main

import "fmt"

func main() {
	//            chave    valor
	usuario := map[string]string{
		"nome":      "John",
		"sobrenome": "Doe",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Mary",
			"ultimo":   "Doe",
		},
		"curso": {
			"nome":   "direito",
			"campus": "central",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "gemeos",
	}

	fmt.Println(usuario2)
}
