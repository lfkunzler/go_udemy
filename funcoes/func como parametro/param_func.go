// go nao eh uma linguagem funcional em essencia
// o fato de poder armazenar funcoes em variaveis permite que possamos
// aplicar os principios da programacao funcional

package main

import "fmt"

func multiplicacao(a, b int) (ret int) {
	ret = a * b
	return
}

func divisao(a, b int) (ret int) {
	ret = a / b
	return
}

func soma(a, b int) (ret int) {
	ret = a + b
	return
}

func subtracao(a, b int) (ret int) {
	ret = a - b
	return
}

func calculadora(f func(int, int) int, p1, p2 int) int {
	return f(p1, p2)
}

func main() {
	fmt.Println("3 x 2 =", calculadora(multiplicacao, 3, 2))
	fmt.Println("3 / 2 =", calculadora(divisao, 3, 2))
	fmt.Println("3 + 2 =", calculadora(soma, 3, 2))
	fmt.Println("3 - 2 =", calculadora(subtracao, 3, 2))
}
