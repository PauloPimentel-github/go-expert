package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Name         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpperCase(name string) string {
	return strings.ToUpper(name)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"ToUpperCase": ToUpperCase,
	})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{Name: "Go Expert", CargaHoraria: 40},
		{Name: "JavaScript Expert", CargaHoraria: 20},
		{Name: "TypeScript Expert", CargaHoraria: 10},
	})
	if err != nil {
		panic(err)
	}
}
