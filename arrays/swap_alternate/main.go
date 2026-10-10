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

	i, j := 0, 1
	for i < n && j < n {
		arr[i], arr[j] = arr[j], arr[i]
		i += 2
		j += 2
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", arr[i])
	}
	fmt.Fprintln(out)
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
