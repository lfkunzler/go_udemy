package main

import "fmt"

func main() {
	pi := 3.14159265

	fmt.Printf("O valor de x é %f\n", pi)
	fmt.Printf("O valor de x é %.4f\n", pi)

	xs := fmt.Sprint(pi)
	fmt.Println("O valor de x é", xs)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", fmt.Sprint(pi))
	fmt.Println("O valor de x é " + fmt.Sprint(pi))

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("%d %f %.1f %t %s", a, b, b, c, d)
}
