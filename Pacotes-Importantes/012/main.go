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
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Name}} - Carga Horária: {{.CargaHoraria}} horas")
	err := tmp.ExecuteTemplate(os.Stdout, "CursoTemplate", curso)
	if err != nil {
		panic(err)
	}
}
