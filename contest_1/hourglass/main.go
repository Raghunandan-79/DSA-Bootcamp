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

	for i := n; i >= 1; i-- {
		for j := 0; j < n - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(".")
		}
		fmt.Println()
	}

	for i := 2; i <= n; i++ {
		for j := 0; j < n - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}
