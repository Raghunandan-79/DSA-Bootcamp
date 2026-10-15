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

	mpp := make(map[int64]int64)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	
		need := k - arr[i]

		if index, exists := mpp[need]; exists {
			fmt.Fprintf(out, "%d %d", index, i + 1)
			return
		}

		mpp[arr[i]] = i + 1
	}

	fmt.Fprintln(out, -1)
}
