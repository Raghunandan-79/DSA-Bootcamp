package main

import "fmt"

func factorial(n int64) int64 {
	fact := int64(1)

	for i := int64(2); i <= n; i++ {
		fact *= i
	}

	return fact
}

func main() {
	var n, r int64
	fmt.Scan(&n, &r)

	n_factorial := factorial(n)
	r_factorial := factorial(r)
	n_minus_r_factorial := factorial(n - r)

	ans := n_factorial / (r_factorial * n_minus_r_factorial)
	fmt.Println(ans)
}
