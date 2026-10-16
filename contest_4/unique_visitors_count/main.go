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

	seen := make(map[int64]bool)
	distinct := int64(0)

	for i := int64(0); i < n; i++ {
		var x int64
		fmt.Fscan(in, &x)

		if !seen[x] {
			seen[x] = true
			distinct++
		}

		fmt.Fprint(out, distinct, " ")
	}
	fmt.Fprintln(out)
}
