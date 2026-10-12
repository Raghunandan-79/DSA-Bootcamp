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

	for column := 0; column < n; column++ {
		for row := n - 1; row >= 0; row-- {
			fmt.Fprint(out, matrix[row][column])
			if row == 0 {
				fmt.Fprintln(out)
			} else {
				fmt.Fprint(out, " ")
			}
		}
	}
}
