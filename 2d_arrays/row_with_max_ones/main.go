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

	var n, m int
	fmt.Fscan(in, &n, &m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	ones_row, max_ones := -1, -1
	for i := 0; i < n; i++ {
		count_ones := 0
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				count_ones++
			}
		}

		if count_ones > max_ones {
			max_ones = count_ones
			ones_row = i
		} else if count_ones == 0 {
			ones_row = -1
		}
	}

	fmt.Fprintln(out, ones_row)
}
