// conversoes de variaveis
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f32 float32 = 6.9
	f64 := 1.1313

	fmt.Println(f32 / float32(f64))

	// cuidado... nao eh tao simples converter um numero
	fmt.Println("Teste", string(97))

	fmt.Println("teste certo", strconv.Itoa(97))

	// o metodo strconv.Atoi() retorna dois valores (pasmem!)
	// o _ ignora o retorno o.o
	conv, _ := strconv.Atoi("123")
	conv2, erro := strconv.Atoi("321")
	fmt.Println(conv)
	fmt.Println(conv2, erro)

	// qualquer coisa diferente de 1, true, TRUE e True sao falsas
	convb, _ := strconv.ParseBool("TRUE")
	if convb {
		fmt.Println("Verdadeiro")
	}
}
