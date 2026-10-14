package main

import (
	"bufio"
	"fmt"
	"os"
)

func digit_sum(n int64) int64 {
	sum := int64(0)

	for n != 0 {
		last_digit := n % 10
		sum += last_digit
		n /= 10
	}

	return sum
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
	if digit_sum(arr[0]) == k {
		prefix[0] = 1
	} else {
		prefix[0] = 0
	}

	for i := int64(1); i < n; i++ {
		if digit_sum(arr[i]) == k {
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
