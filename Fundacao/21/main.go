package main

import (
	"curso-go/Fundacao/21/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)

	carro := matematica.Carro{Marca: "Ford"}
	fmt.Println("Carro: ", carro.Marca)

	carro2 := matematica.Carro{}
	fmt.Println(carro2.Acelerar())
}
