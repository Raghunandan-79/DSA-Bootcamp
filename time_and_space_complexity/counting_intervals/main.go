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

	var q int
	fmt.Fscan(in, &q)

	for q > 0 {
		var t int
		var l, r int64
		fmt.Fscan(in, &t, &l, &r)

		if l > r {
			fmt.Fprintln(out, 0)
			q--
			continue
		}

		count := r - l - 1
		if t == 2 || t == 4 {
			count++
		}

		if t == 3 || t == 4 {
			count++
		}

		if count < 0 {
			count = 0
		}

		fmt.Fprintln(out, count)

		q--
	}
}
