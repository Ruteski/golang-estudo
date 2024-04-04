package main

import (
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
	dogJson := `{"nome":"Rex","raca":"Pitbull","idade":3}`

	var c cachorro
	// ou
	// c := cachorro{}

	//[]byte(dogJson) -> conversao de tipo para slice de byte
	// passado ponteiro para a variavel c para alterar o valor em memoria dela
	if err := json.Unmarshal([]byte(dogJson), &c); err != nil {
		log.Fatal(err)
	}

	// json convertido para struct
	fmt.Println(c)

	// conversao de json para map
	dogJson2 := `{"nome":"Tobby","raca":"Tombalata"}`
	c2 := make(map[string]string)

	if err := json.Unmarshal([]byte(dogJson2), &c2); err != nil {
		log.Fatal(err)
	}

	// json convertido para map
	fmt.Println(c2)

}
