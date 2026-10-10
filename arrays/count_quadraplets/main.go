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

	var n, x int64
	fmt.Fscan(in, &n, &x)

	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	count_quadraplets := 0
	for i := int64(0); i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					if arr[i] - 2 * arr[j] + 3 * arr[k] - 4 * arr[l] == x {
						count_quadraplets++
					}
				}
			}
		}
	}

	fmt.Fprintln(out, count_quadraplets)
}
