package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Name         string
	CargaHoraria int
}

func main() {
	curso := Curso{Name: "Go Expert", CargaHoraria: 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Name}} - Carga Horária: {{.CargaHoraria}} horas"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
