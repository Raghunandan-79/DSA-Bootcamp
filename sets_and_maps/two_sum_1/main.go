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

	set := make(map[int64]struct{})
	for _, val := range arr {
		if _, exists := set[k - val]; exists {
			fmt.Fprintln(out, "TRUE")
			return
		}

		set[val] = struct{}{}
	}

	fmt.Fprintln(out, "FALSE")
}
