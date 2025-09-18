package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	ph := Cliente{
		Nome:  "PH",
		Idade: 30,
		Ativo: true,
	}

	ph.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", ph.Nome, ph.Idade, ph.Ativo)
}
