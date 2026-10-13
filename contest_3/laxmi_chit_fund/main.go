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

	var n int64
	fmt.Fscan(in, &n)

	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	minimum, maximum := arr[0], arr[0]
	total, average := arr[0], int64(0)

	for i := int64(1); i < n; i++ {
		if arr[i] > maximum {
			maximum = arr[i]
		} else if arr[i] < minimum {
			minimum = arr[i]
		}

		total += arr[i]
	}

	average = total / n

	fmt.Fprintf(out, "%d %d %d %d", total, maximum, minimum, average)
}
