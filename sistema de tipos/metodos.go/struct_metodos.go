package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nome_completo string) {
	partes := strings.Split(nome_completo, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Luis", "kunzler"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Perereira")
	fmt.Println(p1.getNomeCompleto())
}
