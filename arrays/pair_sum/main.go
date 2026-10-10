package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	var x int64
	fmt.Fscan(in, &x)

	count_pairs := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] + arr[j] == x {
				count_pairs++
			}
		}
	}
	fmt.Fprintln(out, count_pairs)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		solve(in, out)
		t--
	}
}
