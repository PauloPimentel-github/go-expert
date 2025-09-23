package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Name         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{Name: "Go Expert", CargaHoraria: 40},
		{Name: "JavaScript Expert", CargaHoraria: 20},
		{Name: "TypeScript Expert", CargaHoraria: 10},
	})
	if err != nil {
		panic(err)
	}
}
