package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprime_aprovacao(nota float32) string {
	if nota > 10 || nota < 0 {
		return ("nota invalida")
	} else if nota >= 9 {
		return ("A")
	} else if nota >= 8 {
		return ("B")
	} else if nota >= 7 {
		return ("c")
	} else if nota >= 6 {
		return ("D")
	} else if nota >= 5 {
		return ("E")
	} else {
		return ("F")
	}
}

func numeroaleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10) // r % 10
}

func main() {
	if i := numeroaleatorio(); i >= 6 { // inicializando if like a for
		println("Aprovado")
		fmt.Println(imprime_aprovacao(float32(i)))
	} else {
		println("Reprovado")
		fmt.Println(imprime_aprovacao(float32(i)))
	}
}
