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

	fastest_runner, fastest_runner_idx := arr[0], 0
	for i := 1; i < n; i++ {
		if arr[i] < fastest_runner {
			fastest_runner = arr[i]
			fastest_runner_idx = i
		} else if arr[i] == fastest_runner {
			fastest_runner_idx = i
		}
	}

	fmt.Fprintln(out, fastest_runner_idx + 1)
}
