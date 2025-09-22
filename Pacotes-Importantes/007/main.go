package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParams := r.URL.Query().Get("cep")
	if cepParams == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Aqui você pode adicionar a lógica para buscar o CEP usando o valor de cepParams
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
