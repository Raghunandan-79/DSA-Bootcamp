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

	largest, second_largest := -1, -1
	for i := 0; i < n; i++ {
		if arr[i] > largest {
			second_largest = largest
			largest = arr[i]
		} else if arr[i] > second_largest && arr[i] != largest {
			second_largest = arr[i]
		}
	}

	fmt.Fprintln(out, second_largest)
}
