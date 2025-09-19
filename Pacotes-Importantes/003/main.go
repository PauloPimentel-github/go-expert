package main

import (
	"fmt"
)

func main() {
	// executa a operacao por ultimo
	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
