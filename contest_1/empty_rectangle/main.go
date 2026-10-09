package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i == 1 || j == 1 || i == n || j == m {
				fmt.Print("^")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
