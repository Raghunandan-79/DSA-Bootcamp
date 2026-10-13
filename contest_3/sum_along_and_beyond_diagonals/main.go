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

	total_sum, diagonal_sum := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			total_sum += matrix[i][j]

			if i == j || i + j == n - 1 {
				diagonal_sum += matrix[i][j]
			}
		}
	}

	fmt.Fprintf(out, "%d %d", diagonal_sum, total_sum - diagonal_sum)
}
