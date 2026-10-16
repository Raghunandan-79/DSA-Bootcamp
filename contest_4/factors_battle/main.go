package main

import (
	"bufio"
	"fmt"
	"os"
)

func count_divisors(n int64) int64 {
	count := int64(0)

	for i := int64(1); i * i <= n; i++ {
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

	var a, b int64
	fmt.Fscan(in, &a, &b)

	if count_divisors(a) > count_divisors(b) {
		fmt.Fprintln(out, "A")
	} else if count_divisors(a) < count_divisors(b) {
		fmt.Fprintln(out, "B")
	} else {
		fmt.Fprintln(out, "DRAW")
	}
}
