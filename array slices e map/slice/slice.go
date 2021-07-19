package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{3, 2, 1} // slice: refere-se a uma parte de array

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [...]int{123, 124, 215, 124, 634}
	s2 := a2[1:3] //[elemento inicial, elemento final nao inclusivo]
	// quando o indice inicial for igual a zero, pode ser suprimido
	s3 := a2[:2]
	// tb pode se obter um slice de slice
	s4 := s2[:1]
	fmt.Println(a2, s2, s3, s4)

}
