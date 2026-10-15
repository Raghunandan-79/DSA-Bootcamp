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

	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	mpp := make(map[int64]int64)
	mpp[0] = 0

	pref := int64(0)
	for i := int64(1); i <= n; i++ {
		pref += arr[i - 1]

		if start, ok := mpp[pref - k]; ok {
			fmt.Fprintln(out, start + 1, i)
			return
		}

		if _, ok := mpp[pref]; !ok {
			mpp[pref] = i
		}
	}

	fmt.Fprintln(out, -1)
}
