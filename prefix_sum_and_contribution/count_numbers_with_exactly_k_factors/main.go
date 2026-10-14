package main

import (
	"bufio"
	"fmt"
	"os"
)

func count_factors(n int64) int64 {
	count := int64(0)

	for i := int64(1); i <= n / i; i++ {
		if n % i == 0 {
			count++

			if i != n / i {
				count++
			}
		}
	}

	return count
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q, k int64
	fmt.Fscan(in, &n, &q, &k)

	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	prefix := make([]int64, n)
	if count_factors(arr[0]) == k {
		prefix[0] = 1
	} else {
		prefix[0] = 0
	}

	for i := int64(1); i < n; i++ {
		if count_factors(arr[i]) == k {
			prefix[i] = prefix[i - 1] + 1
		} else {
			prefix[i] = prefix[i - 1]
		}
	}

	for q > 0 {
		var l, r int64
		fmt.Fscan(in, &l, &r)

		l--
		r--

		if l == 0 {
			fmt.Fprintln(out, prefix[r])
		} else {
			fmt.Fprintln(out, prefix[r] - prefix[l - 1])
		}

		q--
	}
}
