package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"PH": 1000, "Ana": 2000, "Lia": 3000}
	m2 := map[string]float64{"PH": 100.20, "Ana": 2000.3, "Lia": 300.0}
	m3 := map[string]MyNumber{"PH": 1000, "Ana": 2000, "Lia": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
