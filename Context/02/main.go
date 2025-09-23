package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Requisição iniciada")
	defer log.Println("Requisição finalizada")
	select {
	case <-time.After(5 * time.Second):
		// Imprime no comand line stout
		log.Println("Requisição processada com sucesso!")
		// Imprime no browser
		w.Write([]byte("Hello, World!"))
	case <-ctx.Done():
		err := ctx.Err()
		// Imprime no comand line stout
		log.Println("Requisição cancelada pelo cliente", err)
		// Imprime no browser
		http.Error(w, "Requisição cancelada pelo cliente", http.StatusRequestTimeout)
	}
}
