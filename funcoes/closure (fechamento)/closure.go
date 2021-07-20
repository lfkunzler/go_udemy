// o conceito de closure (fechamento) diz respeito sobre a funcao interna e seu contexto

package main

import "fmt"

// a funcao retorna uma funcao
func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX() // essa funcao imprime o valor definido na linha 9
	// as funcoes em go tem noção de "suas origens"
	// respeitam o contexto maior em que estão inseridas

}
