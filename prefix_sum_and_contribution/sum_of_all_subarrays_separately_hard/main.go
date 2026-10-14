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

	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	prefix := make([]int64, n+1)
	for i := int64(0); i < n; i++ {
		prefix[i + 1] = prefix[i] + arr[i]
	}

	for i := int64(0); i < n; i++ {
		for j := i; j < n; j++ {
			sum := prefix[j + 1] - prefix[i]
			fmt.Fprintln(out, sum)
		}
	}
}
