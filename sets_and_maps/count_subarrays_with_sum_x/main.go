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
	mpp[0] = 1

	var pref, ans int64
	for i := int64(0); i < n; i++ {
		var x int64
		fmt.Fscan(in, &x)

		pref += x
		ans += mpp[pref - k]
		mpp[pref]++
	}

	fmt.Fprintln(out, ans)
}
