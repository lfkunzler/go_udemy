package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma +", a+b)
	fmt.Println("subtração -", a-b)
	fmt.Println("multiplicacao *", a*b)
	fmt.Println("divisao /", a/b)
	fmt.Println("modulo %", a%b)

	// bitwise
	fmt.Println("or |", a|b)
	fmt.Println("and &", a&b)
	fmt.Println("xor |", a^b)

	// operacoes da lib math
	fmt.Println("maior entre", math.Max(float64(a), float64(b)))
	fmt.Println("menor entre", math.Min(float64(a), float64(b)))
	fmt.Println("potencia entre", math.Pow(float64(a), float64(b)))
	fmt.Println("raiz quad entre", math.Sqrt(float64(a)))
	fmt.Println("potencia entre", math.Pow(float64(a), float64(0.5)))

}
