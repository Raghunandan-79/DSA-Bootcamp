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

	var n, x int
	fmt.Fscan(in, &n, &x)

	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	for i := 0; i < n; i++ {
		if arr[i] == int64(x) {
			fmt.Fprintln(out, "YES")
			return
		}
	}
	
	fmt.Fprintln(out, "NO")
}
