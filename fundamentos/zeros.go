// qual o valor da variavel quando nada eh atribuido?
package main

import "fmt"

func main() {
	var i int
	var f32 float32
	var f64 float64
	var s string
	var b bool
	var p *int // ponteiro pra int

	fmt.Printf("%v %v %v %v %v %v", i, f32, f64, s, b, p)
}
