package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw struct {
}

func (bm bmw) ligarTurbo() {
	fmt.Println("Ligando turbo...")
}

func (bm bmw) fazerBaliza() {
	fmt.Println("Fazendo baliza...")
}

func main() {
	var b esportivoLuxuoso = bmw{}
	b.fazerBaliza()
	b.ligarTurbo()
}
