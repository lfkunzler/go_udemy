package main

import "fmt"

func main() {
	funcsporletra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva":  123125.12,
			"Guga Perereira": 123.14,
		},
		"J": {
			"Joao Silva":     123124.14,
			"Junior Antonio": 12124.4,
		},
		"P": {
			"Pedro Jagunso": 1200.0,
			"Paulo naba":    1212.3,
		},
	}
	for letra, funcionarios := range funcsporletra {
		fmt.Println(letra, funcionarios)
		for nome, salario := range funcionarios {
			fmt.Println(nome, salario)
		}
		fmt.Println()
	}

	fmt.Println("DELETANDOO O P")
	delete(funcsporletra, "P")
	for letra, funcionarios := range funcsporletra {
		fmt.Println(letra, funcionarios)
		for nome, salario := range funcionarios {
			fmt.Println(nome, salario)
		}
		fmt.Println()
	}

	delete(funcsporletra["J"], "Joao Silva")
	fmt.Println("DELETANDOO O JOAO!!!")
	delete(funcsporletra, "P")
	for letra, funcionarios := range funcsporletra {
		fmt.Println(letra, funcionarios)
		for nome, salario := range funcionarios {
			fmt.Println(nome, salario)
		}
		fmt.Println()
	}
}
