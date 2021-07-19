package main

import "fmt"

func main() {
	var s string = "que doidera essa linguagem vei"
	fmt.Println(s)

	i := 3  // por inferencia
	i += 2  // atribui somando
	i -= 2  // subtraindo
	i *= 4  // multiplicando
	i /= 4  // dividindo
	i %= 10 // por resto

	fmt.Println(i)

	x, y := 23, 13 // duas ao mesmo tempo
	fmt.Println(x, y)

	// invertendo sem usar var auxiliar!!!!!!! #eita
	x, y = y, x
	fmt.Println(x, y)
}
