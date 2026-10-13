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
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &grid[i][j])
		}
	}

	ans := make([]int, 0)
	visit := func(value int) bool {
		if value == -1 {
			return false
		}
		ans = append(ans, value)
		return true
	}

	for r := n - 1; r >= 0; r-- {
		if !visit(grid[r][0]) {
			goto output
		}
	}
	for c := 1; c < m; c++ {
		if !visit(grid[0][c]) {
			goto output
		}
	}
	for r := 1; r < n; r++ {
		if !visit(grid[r][m-1]) {
			goto output
		}
	}
	for c := m - 2; c >= 0; c-- {
		if !visit(grid[n-1][c]) {
			goto output
		}
	}

output:
	for i, value := range ans {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, value)
	}
	fmt.Fprintln(out)
}
