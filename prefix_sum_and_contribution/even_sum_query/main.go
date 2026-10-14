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

	var n int64
	fmt.Fscan(in, &n)

	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	prefix := make([]int64, n)
	prefix[0] = 0

	for i := int64(1); i < n; i++ {
		prefix[i] = prefix[i-1]

		if i%2 == 1 {
			prefix[i] += arr[i]
		}
	}

	var q int64
	fmt.Fscan(in, &q)

	for q > 0 {
		var l, r int64
		fmt.Fscan(in, &l, &r)
		l--
		r--

		if l == 0 {
			fmt.Fprintln(out, prefix[r])
		} else {
			fmt.Fprintln(out, prefix[r]-prefix[l-1])
		}

		q--
	}
}
