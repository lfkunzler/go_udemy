package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces sao implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$: %.2f", p.nome, p.preco)
}

func (p produto) valorTotal() float64 {
	return p.preco
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	// cria uma interface imprimivel
	var coisa imprimivel = pessoa{"Roberto", "Pessoa"}
	fmt.Println(coisa.toString())

	imprimir(coisa)

	// polimorfismo de coisa
	// como pessoa e produto possuem o metodo toString()
	// eh possivel associar a interface a ambos durante o runtime
	coisa = produto{"calça jeans", 79.70}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"blusa de lã", 120.99}
	imprimir(p2)
}
