// o defer refere-se ao atraso que pode ser dado a execucao de algum codigo
// por exemplo, se queremos fechar a conexao com o banco de dados ao final de um método
// o defer garante que a conexão será fechada realmente só no final do metodo

package main

import "fmt"

func obtervaloraprovado(n int) int {
	defer fmt.Println("fim!")
	if n > 5000 {
		fmt.Println("valor alto:", n)
		return 5000
	}
	fmt.Println("Valor baixo:", n)
	return n
}

func main() {
	fmt.Println("main:", obtervaloraprovado(23434))
	fmt.Println("main:", obtervaloraprovado(3131))
}
