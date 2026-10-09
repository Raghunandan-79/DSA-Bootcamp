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

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(out, "%d * %d = %d\n", n, i, (n * i))
	}
}
