package main

import "fmt"

func imprime_aprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Println(i+1, "-", aprovado)
	}
}

func main() {
	lista_slice := []string{
		"Maria",
		"Pedro",
		"Rafael",
		"Gabriel",
	}
	imprime_aprovados(lista_slice...)
}
