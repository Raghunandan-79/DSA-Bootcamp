package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	next_non_zero := 0
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[next_non_zero] = arr[i]
			next_non_zero++
		}
	}

	for next_non_zero < n {
		arr[next_non_zero] = 0
		next_non_zero++
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", arr[i])
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
		solve(in , out)
		t--
	}
}
