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

	for j := 0; j < m; j++ {
		fmt.Fprintf(out, "%d ", matrix[0][j])
	}

	for i := 1; i < n; i++ {
		fmt.Fprintf(out, "%d ", matrix[i][m-1])
	}

	if n > 1 {
		for j := m - 2; j >= 0; j-- {
			fmt.Fprintf(out, "%d ", matrix[n-1][j])
		}
	}

	if m > 1 {
		for i := n - 2; i > 0; i-- {
			fmt.Fprintf(out, "%d ", matrix[i][0])
		}
	}
}
