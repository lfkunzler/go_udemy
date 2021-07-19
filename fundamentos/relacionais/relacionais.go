package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("strings: banana == banana? R:", "banana" == "banana")
	fmt.Println("strings: Banana == banana? R:", "Banana" == "banana")

	fmt.Println("Diferente != ", 3 != 2)
	fmt.Println("menor q", 3 < 2)
	fmt.Println("maior q", 3 > 2)
	fmt.Println("menor igual", 3 <= 2)
	fmt.Println("maior igual", 3 >= 2)
	fmt.Println()

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("entre structs com ==", d1 == d2)
	fmt.Println("datas: ", d1.Equal(d2))
}
