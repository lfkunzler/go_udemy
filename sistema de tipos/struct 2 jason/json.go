package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := produto{1, "Notebook", 3999.90, []string{"promocao", "eletronico"}}
	fmt.Println(p1)
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json to struct
	var p2 produto
	jsonstring := `{"id":2,"nome":"Caneta","preco":8.9,"tags":["papelaria","Importado"]}`
	fmt.Println(string(jsonstring))
	json.Unmarshal([]byte(jsonstring), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Tags[1])
}
