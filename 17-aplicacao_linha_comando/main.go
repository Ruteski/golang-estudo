package main

import (
	"fmt"
	"linha-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()

	// err := aplicacao.Run(os.Args) // parametro padrao

	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
