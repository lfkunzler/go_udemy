package main

import "fmt"

func main() {
	// o compilador define o tamanho do array com base na qtde de elementos
	numeros := [...]int{10, 20, 30, 40, 50} // array de inteiros com [5]

	// range retorna 2 numeros, o indice atual e o len do array
	for i, numero := range numeros {
		fmt.Printf("[%d] = %d\n", i, numero)
	}

	// pode-se tambem ignorar o indice, acessando os numeros em ordem
	for _, num := range numeros {
		fmt.Println(num)
	}
}
