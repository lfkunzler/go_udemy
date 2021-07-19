package main

import "fmt"

func main() {
	// se nao definir a capacidade do s1, ha a possibilidade do compilador nao atribuir o mesmo
	// array internamente quando houver a realocacao
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3, 4)

	fmt.Println(s1, s2)
	fmt.Printf("tamanho s1: %d, capacidade s1: %d\n", len(s1), cap(s1))
	fmt.Printf("tamanho s2: %d, capacidade s2: %d\n\n", len(s2), cap(s2))
	// mudando um elemento do slice 1, provamos que no slice 2 acontece o mesmo
	s1[7] = 123
	fmt.Println(s1, s2)
	fmt.Printf("tamanho s1: %d, capacidade s1: %d\n", len(s1), cap(s1))
	fmt.Printf("tamanho s2: %d, capacidade s2: %d\n\n", len(s2), cap(s2))

	s1 = append(s1, 9, 7, 6, 5, 4, 3, 1)
	fmt.Println(s1, s2)
	fmt.Printf("tamanho s1: %d, capacidade s1: %d\n", len(s1), cap(s1))
	fmt.Printf("tamanho s2: %d, capacidade s2: %d\n", len(s2), cap(s2))

}
