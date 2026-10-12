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

	var x int
	fmt.Fscan(in, &x)

	for i := 0; i < n; i++ {
		if arr[i] == x {
			fmt.Fprintln(out, "YES")
			return
		}
	}

	fmt.Fprintln(out, "NO")
}
