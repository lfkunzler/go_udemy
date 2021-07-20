// operador para funcao variatica ...

package main

import "fmt"

func media(numeros ...float64) (ret float64) {
	ret = 0

	for _, num := range numeros {
		ret += num
	}
	ret /= float64(len(numeros))
	return
}

func main() {
	fmt.Println("media:", media(5, 7, 5, 7))
}
