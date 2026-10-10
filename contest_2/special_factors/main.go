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

	var n int
	fmt.Fscan(in, &n)

	found := false
	for i := 1; i <= n; i++ {
		if n % i == 0 && (i % 10 == 2 || i % 10 == 7) {
			fmt.Fprintf(out, "%d ", i)
			found = true
		}
	}

	if !found {
		fmt.Fprintln(out, -1)
	}
}
