package main

import "fmt"

func recuperaExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execução")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperaExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA E EXATAMENTE 6!")
}

func main() {

	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")

}
