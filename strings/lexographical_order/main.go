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

	var a, b string
	fmt.Fscan(in, &a, &b)

	if a < b {
		fmt.Fprintln(out, "A")
	} else if a > b {
		fmt.Fprintln(out, "B")
	} else {
		fmt.Fprintln(out, "Equal")
	}
}
