package main

import "fmt"

type item struct {
	id    int
	qtde  float64
	preco float64
}

type pedido struct {
	id    int
	itens []item
}

func (p pedido) valorPedido() (total float64) {
	for _, item := range p.itens {
		total += item.preco * item.qtde
	}
	return
}

func main() {
	p1 := pedido{
		id: 1,
		itens: []item{
			{1, 2, 10.4},
			{2, 4, 1100.6},
			{11, 100, 3.3},
		},
	}
	fmt.Println(p1)
	fmt.Printf("Valor total do pedido é de R$%.2f", p1.valorPedido())
}
