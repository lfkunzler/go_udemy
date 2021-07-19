// funcoes sao cidadãos de primeira classe e podem ser atribuidas a variaveis
package main

import "fmt"

// cria uma variavel que recebe o resultado de uma func anonima
var soma = func(a, b int) int {
	return a + b
}

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func calculate(fp func(int, int) int) {
	ans := fp(3, 2)
	fmt.Printf("\n%v\n", ans)
}

func main() {
	fmt.Println(soma(1, 2))

	// funcao anonima local
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(3, 2))

	// funcoes que recebem funcao como parametro
	calculate(Plus)
	calculate(Minus)
	calculate(Multiply)
}
