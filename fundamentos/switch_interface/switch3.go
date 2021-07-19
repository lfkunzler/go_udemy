package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "nao sei"
	}
}

func main() {
	var u8 uint8 = 10
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("oi"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(u8))
	fmt.Println(tipo(time.Now()))

}
