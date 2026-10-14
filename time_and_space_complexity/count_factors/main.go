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

	cnt := 0
	for i := int64(1); i * i <= n; i++ {
		if n % i == 0 {
			cnt++

			if i != n / i {
				cnt++
			}
		}
	}

	fmt.Fprintln(out, cnt)
}
