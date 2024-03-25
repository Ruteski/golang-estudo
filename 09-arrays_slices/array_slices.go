package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "posicao 1"
	fmt.Println(array2)

	array3 := [5]string{"posicao 1", "posicao 2", "posicao 3", "posicao 4", "posicao 5"}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array4)

	// SLICE
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array4))

	slice = append(slice, 35)
	fmt.Println(slice)

	slice2 := array3[1:3] // ponteiro
	fmt.Println(slice2)

	array3[1] = "Posicao alterada"
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	fmt.Println("----------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacidade

}
