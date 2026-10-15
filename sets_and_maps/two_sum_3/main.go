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

	var n, k int64
	fmt.Fscan(in, &n, &k)

	mpp := make(map[int64]int64)
	ans := int64(0)

	for i := int64(0); i < n; i++ {
		var x int64
		fmt.Fscan(in, &x)
		ans += mpp[k - x]
		mpp[x]++
	}

	fmt.Fprintln(out, ans)
}
