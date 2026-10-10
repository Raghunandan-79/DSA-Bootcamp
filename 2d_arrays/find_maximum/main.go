package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	maximum := math.MinInt
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] > maximum {
				maximum = matrix[i][j]
			}
		}
	}

	fmt.Println(maximum)
}
