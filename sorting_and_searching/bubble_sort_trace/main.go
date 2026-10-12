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

	for pass := 1; pass < n; pass++ {
		swaps := 0

		for i := 0; i < n - pass; i++ {
			if arr[i] > arr[i + 1] {
				arr[i], arr[i + 1] = arr[i + 1], arr[i]
				swaps++
			}
		}

		fmt.Fprintf(out, "Pass %d: ", pass)
		for i := 0; i < n; i++ {
			if i > 0 {
				fmt.Fprintf(out, " ")
			}
			fmt.Fprint(out, arr[i])
		}

		fmt.Fprintf(out, " , swaps = %d\n", swaps)

		if swaps == 0 {
			break
		}
	}
}
