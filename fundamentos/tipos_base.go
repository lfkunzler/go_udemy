package main

import (
	"fmt"
	"reflect"
)

func main() {

	// pacote reflect diz o tipo da variavel passada
	fmt.Println("literal inteiro é:", reflect.TypeOf(321))

	// sem sinal: uint8=byte uint16 uint32 uint64
	var a byte = 255
	fmt.Println("byte eh", reflect.TypeOf(a))

	// nao tem char (vish)
	c := 'a'
	fmt.Println("char eh", reflect.TypeOf(c))
	fmt.Println(c)

	var f float32 = 123.123
	f_lit := 123.123
	fmt.Println("float declarado 32", reflect.TypeOf(f))
	fmt.Println("float literal", reflect.TypeOf(f_lit)) // float64

	var b bool = true
	b1 := false
	fmt.Println("boolean declarado", reflect.TypeOf(b))
	fmt.Println("boolean literal", reflect.TypeOf(b1))

	s1 := "ola meu nome eh felipe"
	fmt.Println("string declarado", reflect.TypeOf(s1))
	fmt.Println(s1)
	fmt.Println("tamanho dela:", len(s1)) // exatamente quant de char

	s2 := `ola
meu
nome
eh
felipe`
	fmt.Println("string declarado", reflect.TypeOf(s2))
	fmt.Println(s2)
	fmt.Println("tamanho dela:", len(s2)) // exatamente quant de char
}
