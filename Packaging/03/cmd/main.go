package main

import (
	"fmt"

	"github.com/PauloPimentel-github/go-expert/03/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m)
	// fmt.Println(m.Add())
	// fmt.Println("Hello, world!")
}
