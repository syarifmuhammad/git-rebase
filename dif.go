package main

func dif(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
