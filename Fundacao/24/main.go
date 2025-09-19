package main

func main() {
	a := 10
	b := 2
	c := 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a > b {
		println("a é maior que b")
	} else if a < b {
		println("a é menor que b")
	} else {
		println("a é igual a b")
	}

	if a > b && a > c {
		println("a é o maior")
	}

	switch a {
	case 1:
		println("a é 1")
	case 2:
		println("a é 2")
	case 3:
		println("a é 3")
	default:
		println("a não é 1, 2 ou 3")
	}
}
