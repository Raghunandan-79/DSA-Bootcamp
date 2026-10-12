package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	var d int
	fmt.Fscan(in, &d)

	if n == 0 {
		fmt.Fprintln(out)
		return
	}

	d %= n

	slices.Reverse(arr[:d])
	slices.Reverse(arr[d:])
	slices.Reverse(arr)

	for _, value := range arr {
		fmt.Fprint(out, value, " ")
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
