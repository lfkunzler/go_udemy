package main

import "fmt"

func fatorial(a int) int {
	if a > 0 {
		return a * fatorial(a-1)
	}
	return 1
}

func main() {
	fmt.Printf("fatorial de %d é %d", 5, fatorial(5))
}
