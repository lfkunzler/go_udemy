// embora go seja fortemente tipada
// as interfaces possuem polimorfismo...
package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}
	var coisa2 dinamico = "opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"GO (GoLang): Explorando a linguagem do google"}
	fmt.Println(coisa2)
}
