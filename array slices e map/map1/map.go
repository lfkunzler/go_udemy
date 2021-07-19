package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[3655541031] = "felipe"
	aprovados[4503768000] = "jaque"
	aprovados[12124912124] = "unknow"

	for cpf, nome := range aprovados {
		fmt.Printf("Nome: %s CPF: %d\n", nome, cpf)
	}

	fmt.Println(aprovados)
	delete(aprovados, 3655541031)
	fmt.Println(aprovados)
}
