package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Name         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{Name: "Go Expert", CargaHoraria: 40},
			{Name: "JavaScript Expert", CargaHoraria: 20},
			{Name: "TypeScript Expert", CargaHoraria: 10},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)
}
