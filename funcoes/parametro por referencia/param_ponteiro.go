package main

import "fmt"

func inc1(n int) {
	n += 10
}

func inc2(n *int) {
	*n += 10
}

func main() {
	n := 1

	inc1(n)
	fmt.Println("funcao inc 1", n)

	inc2(&n)
	fmt.Println("funcao inc 2", n)
}
