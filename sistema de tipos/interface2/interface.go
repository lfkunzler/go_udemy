package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo      string
	velocidade  int
	turboLigado bool
}

// para a struct obedecer a interface, precisa do metodo ligarTurbo()
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F430", 0, false}
	fmt.Println(carro1)
	carro1.ligarTurbo()
	fmt.Println(carro1)

	// quando a interface altera o conteudo da variavel
	// eh necessario passar a variavel por ponteiro para a interface
	var carro2 esportivo = &ferrari{"F40", 0, false}
	fmt.Println(carro2)
	carro2.ligarTurbo()
	fmt.Println(carro2)

}
