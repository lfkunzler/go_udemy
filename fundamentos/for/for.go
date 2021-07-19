package main

import (
	"fmt"
	"time"
)

func main() {
	// nao existe while em go #F
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println("par")
		} else {
			fmt.Println("impar")
		}
	}

	for { // infinito
		fmt.Println("loop infinito", time.Second)
		time.Sleep(time.Second * 5)
	}
}
