package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//escrita
	tamanho, err := f.Write([]byte("Escrevendo no arquivo\n"))
	// tamanho, err = f.WriteString("Escrevendo mais uma linha\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com %d bytes\n", tamanho)
	f.Close()

	//leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Printf("Lidos %d bytes: %s\n", n, string(buffer[:n]))
	}
	arquivo2.Close()

	//remover arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
