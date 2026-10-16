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

	total := int64(0)
	for _, value := range arr {
		total += value
	}

	count := int64(0)
	left := int64(0)
	for _, value := range arr {
		left += value
		if left == total - left + value {
			count++
		}
	}

	fmt.Fprintln(out, count)
}
