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

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	var p int
	fmt.Fscan(in, &p)

	pass, fail := 0, 0
	for i := 0; i < n; i++ {
		if arr[i] >= p {
			pass++
		} else {
			fail++
		}
	}

	fmt.Fprintf(out, "Pass: %d\nFail: %d\n", pass, fail)
}
