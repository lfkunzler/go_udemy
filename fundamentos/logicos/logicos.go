package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	tv50 := trab1 && trab2
	tv32 := trab1 != trab2 // xor logico (n existe)
	sorvete := trab1 || trab2

	return tv50, tv32, sorvete
}

func main() {
	rtv50, rtv32, rsorv := compras(true, false)

	fmt.Printf("tv50: %t | tv32: %t | sorvete: %t | saude: %t.", rtv50, rtv32, rsorv, !rsorv)
}
