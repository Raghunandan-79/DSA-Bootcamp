package main

import "fmt"

func factorial(n int) int {
	fact := 1

	for i := 2; i <= n; i++ {
		fact *= i
	}

	return fact
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(factorial(n))
}
