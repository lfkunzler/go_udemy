package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type ferrari struct {
	carro       // pseudo heranca
	turboLigado bool
}

func main() {
	f := ferrari{}

	f.nome = "f430"
	f.velocidade = 220
	f.turboLigado = true

	fmt.Println("A ferrari esta com turbo ligado?", f.nome, f.turboLigado)
	fmt.Println(f)
}
