package main

import (
	"fmt"
	"math/rand"
	"time"
)

func nota2conceito(n float32) string {
	nota := int(n)
	ret := "F"

	switch nota { // o
	case 10:
		fallthrough // entra no case abaixo se entrar nesse
	case 9:
		ret = "A"
	case 8:
		ret = "B"
	case 7, 6:
		ret = "C"
	case 5:
		ret = "D"
	case 4:
		ret = "E"
	case 3:
		fallthrough
	case 2, 1, 0:
		ret = "F"
	default:
		ret = "nota invalida"
	}
	return ret
}

func nota2conceito2(nota float32) string {
	switch { // switch doidao, case avalia condicao
	case nota > 10 || nota < 0:
		return ("nota invalida")
	case nota >= 9:
		return ("A")
	case nota >= 8:
		return ("B")
	case nota >= 7:
		return ("C")
	case nota >= 6:
		return ("D")
	case nota >= 5:
		return ("E")
	default:
		return ("F")
	}
}

func numeroaleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(100) // r % 10
}

func main() {
	a := numeroaleatorio()
	var n float32
	n = float32(a) / 10
	fmt.Println(nota2conceito2(n), n)

}
