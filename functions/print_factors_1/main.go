package main

import (
	"bufio"
	"fmt"
	"os"
)

func print_factors(n int) {
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	print_factors(n)
}
