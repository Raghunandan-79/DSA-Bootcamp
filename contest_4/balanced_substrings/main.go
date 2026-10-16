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
	var s string
	fmt.Fscan(in, &n, &q)
	fmt.Fscan(in, &s)

	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] - 1
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			prefix[i + 1] += 2
		}
	}

	for ; q > 0; q-- {
		var l, r int
		fmt.Fscan(in, &l, &r)
		if prefix[r] == prefix[l - 1] {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
