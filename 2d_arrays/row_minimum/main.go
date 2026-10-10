package main

import (
	"bufio"
	"fmt"
	"math"
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

	for i := 0; i < n; i++ {
		minimum := math.MaxInt
		for j := 0; j < m; j++ {
			if matrix[i][j] < minimum {
				minimum = matrix[i][j]
			}
		}
		fmt.Fprintf(out, "%d ", minimum)
	}
}
