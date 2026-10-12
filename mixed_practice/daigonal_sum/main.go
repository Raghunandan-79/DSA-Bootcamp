package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				sum += matrix[i][j]
			}
		}
	}
	fmt.Fprintln(out, sum)
}
