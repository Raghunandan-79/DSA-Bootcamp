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

	seen := make(map[int64]struct{})
	seen[0] = struct{}{}

	var prefix int64 = 0
	for _, val := range arr {
		prefix += val

		if _, exists := seen[prefix - k]; exists {
			fmt.Fprintln(out, "YES")
			return
		} 

		seen[prefix] = struct{}{}
	}

	fmt.Fprintln(out, "NO")
}
