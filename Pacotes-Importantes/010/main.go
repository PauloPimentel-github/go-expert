package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	/*mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})*/
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{titile: "Meu blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type blog struct {
	titile string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.titile))
}
