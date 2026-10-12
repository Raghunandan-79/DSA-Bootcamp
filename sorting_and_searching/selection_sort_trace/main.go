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

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	for i := 0; i < n - 1; i++ {
		min_idx := i
		min_val := arr[i]

		for j := i + 1; j < n; j++ {
			if arr[j] < min_val {
				min_val = arr[j]
				min_idx = j
			}
		}

		arr[i], arr[min_idx] = arr[min_idx], arr[i]

		fmt.Fprintf(out, "Pass %d: ", i + 1)
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, "%d", arr[j])
		}
		fmt.Fprintf(out, " , min_selected = %d\n", min_val)
	}	
}
