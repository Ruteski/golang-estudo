package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     "string",
	}

	fmt.Println(mapa)
}
