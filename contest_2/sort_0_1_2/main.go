package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	slices.SortFunc(arr, func(a, b int) int {
		return b - a
	})

	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", arr[i])
	}
	fmt.Fprintln(out)
}
