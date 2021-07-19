package main

import (
	"fmt"
	"reflect"
)

func main() {
	// cria um slice de 10 elementos
	// a capacidade maxima do slice (para realocacao) eh 20 elementos
	// ja ha reservada na memoria
	s1 := make([]int, 10, 20)

	s1[9] = 10

	fmt.Println(s1, reflect.TypeOf(s1))
	fmt.Printf("tamanho: %d, capacidade %d\n", len(s1), cap(s1))

	// anexa ao fim do slice (posicao 10) mais 10 elementos
	s1 = append(s1, 123, 124, 64, 35, 234, 56, 345, 142, 75, 45)

	fmt.Println(s1, reflect.TypeOf(s1))
	fmt.Printf("tamanho: %d, capacidade %d\n", len(s1), cap(s1))

	// se exceder a capacidade interna do array referenciado pelo slice
	// um novo array eh alocado para aumentar a capacidade
	s1 = append(s1, 12)
	fmt.Println(s1, reflect.TypeOf(s1))
	fmt.Printf("tamanho: %d, capacidade %d\n", len(s1), cap(s1))
}
