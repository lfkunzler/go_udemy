package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3}
	var s1 []int

	s1 = append(s1, 4, 5, 6)
	fmt.Println(a1, s1)

	s2 := make([]int, 2)
	copy(s2, s1) // so copia dois elementos, pq eh a capacidade do s2
	fmt.Println(a1, s1, s2)

	fmt.Printf("%p\n", &s1)
}
