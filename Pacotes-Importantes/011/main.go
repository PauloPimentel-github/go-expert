package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// 1. Registre a rota mais específica primeiro.
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blog do Paulo"))
	})
	// 2. Crie um FileServer para servir o diretório estático.
	fileServer := http.FileServer(http.Dir("./public"))
	// 3. Registre a rota que serve o arquivo.
	mux.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
