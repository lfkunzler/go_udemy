package main

import "fmt"

func main() {
	// homogenea (elementos do mesmo tipo) e estatico (nao muda de tamanho)
	var notas [3]float32
	var media float32

	notas[0], notas[1], notas[2] = 6, 10, 4

	for i := 0; i < len(notas); i++ {
		media += notas[i]
	}
	media /= float32(len(notas))

	fmt.Println(notas)
	fmt.Println(media)
}
