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

	var q int64
	fmt.Fscan(in, &q)

	for q > 0 {
		var l, r int64
		fmt.Fscan(in, &l, &r)
		
		fmt.Fprintln(out, (r - l + 1) * (l + r) / 2)

		q--
	}
}
