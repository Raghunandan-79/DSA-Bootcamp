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

	low, high := 0, n - 1
	for low <= high {
		mid := low + (high - low) / 2

		if arr[mid] == x {
			fmt.Fprintln(out, "YES")
			return
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Fprintln(out, "NO")
}
