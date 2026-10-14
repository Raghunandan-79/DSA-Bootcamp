package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int64
	fmt.Fscan(in, &n, &k)

	factors := []int64{}
	for i := int64(1); i <= n/i; i++ {
		if n%i == 0 {
			factors = append(factors, i)

			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}

	slices.Sort(factors)

	if k > int64(len(factors)) {
		fmt.Fprintln(out, -1)
		return
	}

	fmt.Fprintln(out, factors[k-1])
}
