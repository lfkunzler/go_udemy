// a funcao init eh executada quando o codigo eh executado pela primeira vez

package main

import "fmt"

func init() {
	fmt.Println("Codigo iniciado")
}

func main() {
	fmt.Println("a funcao init nao foi chamada, ein!")
}
