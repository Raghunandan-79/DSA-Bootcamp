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

	for i := int64(1); i <= q; i++ {
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

			if mpp[x] > 0 {
				mpp[x]--
			
				if mpp[x] == 0 {
					delete(mpp, x)
				}
			}

		case 3:
			fmt.Fprintln(out, len(mpp))

		case 4:
			var x int64
			fmt.Fscan(in, &x)

			if mpp[x] > 0 {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		}	
	}
}
