package main

import (
	"bufio"
	"fmt"
	"os"
)

func is_prime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	if is_prime(n) {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}
