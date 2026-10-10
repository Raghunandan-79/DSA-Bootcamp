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

	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			for j := 0; j < m; j++ {
				fmt.Printf("%d ", matrix[i][j])
			}
		} else {
			for j := m - 1; j >= 0; j-- {
				fmt.Printf("%d ", matrix[i][j])
			}
		}
	}
}
