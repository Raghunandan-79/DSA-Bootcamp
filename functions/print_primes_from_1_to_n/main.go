package main

import "fmt"

func is_prime(n int64) bool {
	if n < 2 {
		return false
	}

	for i := int64(2); i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func print_primes(n int64) {
	for i := int64(2); i <= n; i++ {
		if is_prime(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	var n int64
	fmt.Scan(&n)

	print_primes(n)
}
