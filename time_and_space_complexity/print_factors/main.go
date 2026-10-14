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

	var n int64
	fmt.Fscan(in, &n)

	ans := []int64{}

	for i := int64(1); i <= n / i; i++ {
		if n % i == 0 {
			ans = append(ans, i)

			if i != n / i {
				ans = append(ans, n / i)
			}
		}
	}

	slices.Sort(ans)

	for i := 0; i < len(ans); i++ {
		fmt.Fprintf(out, "%d ", ans[i])
	}
}
