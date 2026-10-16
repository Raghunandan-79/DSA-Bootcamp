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

	var n, q int
	fmt.Fscan(in, &n, &q)

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var value int
		fmt.Fscan(in, &value)
		prefix[i] = prefix[i-1] + value
	}

	for ; q > 0; q-- {
		var l, r int
		fmt.Fscan(in, &l, &r)
		remaining := prefix[n] - (prefix[r] - prefix[l - 1])
		if remaining % 2 == 1 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
