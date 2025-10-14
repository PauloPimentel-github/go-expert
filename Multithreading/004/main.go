package main

import (
	"fmt"
)

// Thread 1
func main() {
	channel := make(chan string) // vazio

	// Thread 2
	go func() {
		channel <- "Olá Mundo!" // Está cheio
	}()

	// Thread 1
	msg := <-channel //canal esvazia
	fmt.Println(msg)
}
