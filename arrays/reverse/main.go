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

	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	i, j := 0, n - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", arr[i])
	}
}
