package main

import "fmt"

func main() {
	i := 1

	// nao ha aritimetica de ponteiro #f
	var p *int = nil

	p = &i
	fmt.Println(i)
	fmt.Println(p)
	*p++ // dereferencia (incrementa o valor)
	fmt.Println(i)
}
