package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	num := 1
	for i := 0; i < (n + 1) / 2; i++ {
		fmt.Fprintf(out, "%d ", num)
		num += 2
	}

	num = n
	if num % 2 != 0 {
		num--
	}

	for i := 0; i < n / 2; i++ {
		fmt.Fprintf(out, "%d ", num)
		num -= 2
	}
	fmt.Fprintln(out)
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
