package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	json := bytes.NewBuffer([]byte(`{"curso":"Go Expert","nivel":"Avançado","carga_horaria":40}`))
	resp, error := c.Post("https://www.google.com", "application/json", json)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
