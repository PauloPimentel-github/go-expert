package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Microsecond}
	resp, error := c.Get("https://www.google.com")
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(body))
}
