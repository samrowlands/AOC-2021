package main

func main() {
	println(factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	value := n * factorial(n-1)
	return value
}
