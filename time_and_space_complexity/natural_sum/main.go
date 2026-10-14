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

	ans := n * (n + 1) / 2
	fmt.Fprintln(out, ans)
}
