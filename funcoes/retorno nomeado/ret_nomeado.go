package main

import "fmt"

func f1(i1, i2 int) (r1, r2 int) {
	r1 = i2
	r2 = i1
	return // retorno limpo, pq ja atribui as variaveis
}

func main() {
	i1 := 10
	i2 := 20

	fmt.Println(i1, i2)
	i1, i2 = f1(i1, i2)
	fmt.Println(i1, i2)

}
