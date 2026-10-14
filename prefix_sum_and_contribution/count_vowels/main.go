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

	var s string
	fmt.Fscan(in, &s)

	prefix := make([]int, n)
	if s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' {
		prefix[0] = 1
	} else {
		prefix[0] = 0
	}

	for i := int64(1); i < n; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			prefix[i] = prefix[i-1] + 1
		} else {
			prefix[i] = prefix[i-1]
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
