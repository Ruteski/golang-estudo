package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// funcao com varios retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

// pode ter uma funcao que recebe como paramatro outra funcao
// e pode retornar uma funcao tb

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da funcao f")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// ignora o segundo retorno
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)

	_, resultadoSubtracao2 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao2)
}
