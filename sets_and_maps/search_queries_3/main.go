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

	mpp := make(map[int64]int64)
	for i := int64(0); i < n; i++ {
		mpp[arr[i]] = i + 1
	}

	for i := int64(1); i <= q; i++ {
		var x int64
		fmt.Fscan(in, &x)

		_, exists := mpp[x]
		if exists {
			fmt.Fprintln(out, mpp[x])
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
