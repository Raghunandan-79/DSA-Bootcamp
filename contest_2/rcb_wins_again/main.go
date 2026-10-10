package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(arr []int, n int, out *bufio.Writer) {
	result := []int{}

	left := n / 2 - 1
	right := n / 2

	result = append(result, arr[left])
	result = append(result, arr[right])

	left--
	right++

	for left >= 0 && right < n {
		result = append(result, arr[left])
		result = append(result, arr[right])
		left--
		right++
	}

	for i := 0; i < len(result); i++ {
		if i > 0 {
			fmt.Fprintf(out, " ")
		}
		fmt.Fprintf(out, "%d", result[i])
	}
	fmt.Fprintln(out)
}

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

	solve(arr, n, out)
}
