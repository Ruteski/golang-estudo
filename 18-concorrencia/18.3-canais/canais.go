package main

import (
	"fmt"
	"time"
)

func main() {
	// crio um canal de string
	canal := make(chan string)

	go escrever("Ola mundo!", canal)

	fmt.Println("Depois da funcao escrever começar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
