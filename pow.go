package main

func Pow(a, b int) int {
	if b < 0 {
		panic("negative exponent")
	}
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
