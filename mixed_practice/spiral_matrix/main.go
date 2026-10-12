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

	top, bottom := 0, n-1
	left, right := 0, m-1
	spiral := []int{}

	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			spiral = append(spiral, matrix[top][col])
		}
		top++

		for row := top; row <= bottom; row++ {
			spiral = append(spiral, matrix[row][right])
		}
		right--

		if top <= bottom {
			for col := right; col >= left; col-- {
				spiral = append(spiral, matrix[bottom][col])
			}
			bottom--
		}

		if left <= right {
			for row := bottom; row >= top; row-- {
				spiral = append(spiral, matrix[row][left])
			}
			left++
		}
	}

	for _, val := range spiral {
		fmt.Fprintf(out, "%d ", val)
	}
}
