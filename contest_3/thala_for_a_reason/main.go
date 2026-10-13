package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var r int64
	fmt.Fscan(in, &r)

	sum := int64(0)
	for r != 0 {
		last_digit := r % 10
		sum += last_digit
		r /= 10
	}

	if sum == 7 {
		fmt.Fprintln(out, "Thala for a reason")
	} else {
		fmt.Fprintln(out, "Blocked for no reason")
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for t > 0 {
		solve(in, out)
		t--
	}
}
