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

	mpp := make(map[int64]int64)

	for i := int64(0); i < q; i++ {
		var t int64
		fmt.Fscan(in, &t)

		switch t {
		case 1:
			var x int64
			fmt.Fscan(in, &x)

			mpp[x]++

		case 2:
			var x int64
			fmt.Fscan(in, &x)
			delete(mpp, x)

		case 3:
			fmt.Fprintln(out, len(mpp))

		case 4:
			var x int64
			fmt.Fscan(in, &x)

			if _, ok := mpp[x]; ok {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		}
	}
}
