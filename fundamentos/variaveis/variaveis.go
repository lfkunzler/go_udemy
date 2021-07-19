package main

import io "fmt"

func main() {
	const PI float32 = 3.1416
	var raio = 3.2                             // nao precisa dar o tipo quando atribui
	area := PI * float32(raio) * float32(raio) // := cria e ja atribui a variavel

	io.Printf("valor de area: %f\n", area)

	const ( // outra forma de criar constantes
		a = 1
		b = 2
	)

	// var <name<1, 2, ... n>> <type> =
	var e, f bool = true, false

	io.Println(e, f)

	g, h, i := 2, false, "ooi"

	io.Println(g, h, i)
}
