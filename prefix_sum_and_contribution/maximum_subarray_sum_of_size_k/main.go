package main

import (
	"bufio"
	"fmt"
	"math"
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

	prefix := make([]int64, n + 1)
	for i := int64(0); i < n; i++ {
		prefix[i + 1] = prefix[i] + arr[i]
	}

	max_subarray_sum := int64(math.MinInt64)
	
	for i := int64(0); i + k <= n; i++ {
		sum := prefix[i + k] - prefix[i]
		max_subarray_sum = int64(math.Max(float64(sum), float64(max_subarray_sum)))
	}

	fmt.Fprintln(out, max_subarray_sum)
}
