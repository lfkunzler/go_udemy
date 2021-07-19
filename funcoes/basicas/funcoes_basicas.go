package main

import "fmt"

func f1() { // sem parametro, sem retorno
	fmt.Println("F1")
}

func f2(s1 string, s2 string) { // dois parametros, sem retorno
	fmt.Printf("F2: %s %s\n", s1, s2)
}

func f3() string { // sem parametro, retorno string
	return "F3"
}

func f4(s1, s2 string) string { // recebe dois parametros, retorna 1
	return fmt.Sprintf("F4: %s %s", s1, s2)
}

func f5() (string, string) { // nao recebe parametros, retorna duas strings
	return "retorno 1", "retorno 2"
}

func main() {
	f1()
	f2("parametro 1", "parm 2")

	r3, r4 := f3(), f4("p1", "p2")
	fmt.Println(r3, r4)

	r51, r52 := f5()
	fmt.Println(r51, r52)
}
