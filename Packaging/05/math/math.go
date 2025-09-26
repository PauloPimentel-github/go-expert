package math

type math struct {
	a int
	b int
	C int
}

func NewMath(a, b int) math {
	return math{a: a, b: b}
}

func (math math) Add() int {
	return math.a + math.b
}
