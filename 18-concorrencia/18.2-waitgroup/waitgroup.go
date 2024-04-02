package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// faz com que o programa espere o termino das goroutines para entao ser finalizado
	var waitGroup sync.WaitGroup

	//informa o numero de goroutines que ele tem que esperar
	waitGroup.Add(4)

	// funcao anonima com goroutine
	go func() {
		escrever("Ola mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 3!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
