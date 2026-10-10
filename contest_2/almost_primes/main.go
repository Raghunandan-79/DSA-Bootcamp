package main

import (
	"bufio"
	"fmt"
	"os"
)

func count_divisors(n int) int {
	cnt := 0
	
	for i := 1; i * i <= n; i++ {
		if n % i == 0 {
			cnt++

			if i != n / i {
				cnt++
			}
		}
	}

	return cnt
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		if count_divisors(i) <= 4 {
			fmt.Fprintf(out, "%d ", i)
		}
	}
}
