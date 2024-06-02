package main

func Suma(x, y int) int {
	return x + y
}

// Se utiliza go test -cover para saber cuanto porcentaje se esta cumpliendo
func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonnaci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonnaci(n-1) + Fibonnaci(n-2)
}
