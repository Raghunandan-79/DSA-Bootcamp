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

	var n, q int64
	fmt.Fscan(in, &n, &q)

	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	prefix := make([]int64, n)
	prefix[0] = arr[0]

	for i := int64(1); i < n; i++ {
		var val int64

		if i % 2 == 0 {
			val = arr[i]
		} else {
			val = -arr[i]
		}

		prefix[i] = prefix[i - 1] + val
	}

	for q > 0 {
		var l, r int64
		fmt.Fscan(in, &l, &r)
		l--
		r--

		ans := prefix[r]
		if l > 0 {
			ans -= prefix[l - 1]
		}

		if l % 2 == 1 {
			ans = -ans
		}

		fmt.Fprintln(out, ans)

		q--
	}
}
