package main

import (
	"bufio"
	"fmt"
	"os"
)

func is_prime(n int64) bool {
	if n <= 1 {
		return false
	}

	for i := int64(2); i <= n / i; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int64
	fmt.Fscan(in, &n)

	count := 0
	for i := int64(2); i <= n; i++ {
		if is_prime(i) {
			count++
		}
	}

	fmt.Fprintln(out, count)
}
