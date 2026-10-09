package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= 2 * n - 1; i++ {
		k := i
		if k > n {
			k = 2 * n - i
		}

		for j := 1; j < k; j++ {
			fmt.Print(" ")
		}

		fmt.Print(">")
		if k > 1 {
			for j := 1; j <= 2 * k - 3; j++ {
				fmt.Print(" ")
			}
			fmt.Print(">")
		}
		fmt.Println()
	}
}
