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
	
	if n < 2 {
		fmt.Fprintln(out, "NO")
		return
	}

	for i := int64(2); i <= n / i; i++ {
		if n % i == 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
}
