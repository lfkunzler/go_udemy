package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"felipe": 123124.12,
		"luis":   123124.41,
		"jaque":  12314.23,
		"olivia": 132124.14,
		"jaine":  512.12,
	}

	// o primeiro valor do for sempre sera a chave do map
	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	fmt.Printf("\n adicionando rafael \n")
	// adicionando depois
	salarios["rafael"] = 1350.0

	fmt.Printf("\n removendo inexistente \n")
	delete(salarios, "inexistente")

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	fmt.Printf("\n removendo luis \n")
	delete(salarios, "luis")

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

}
