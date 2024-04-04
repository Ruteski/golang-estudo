package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Pitbull", 3}
	fmt.Println(c)

	dogJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dogJson)                  //mostra um slice de bytes
	fmt.Println(bytes.NewBuffer(dogJson)) //mostra o json propriamente dito

	c2 := map[string]string{
		"nome": "Tobby",
		"raca": "Tombalata",
	}

	dogJson2, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dogJson2)
	fmt.Println(bytes.NewBuffer(dogJson2))

}
