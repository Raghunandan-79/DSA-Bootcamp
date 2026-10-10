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

	var a, b int64
	fmt.Fscan(in, &a, &b)

	if a == 0 && b == 1 {
		fmt.Fprintln(out, "Yes")
	} else if a == 1 && b == 0 {
		fmt.Fprintln(out, "Yes")
	}else if a == b {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
