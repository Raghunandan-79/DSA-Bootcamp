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

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		shifts := 0

		for j >= 0 && arr[j] > key {
			arr[j + 1] = arr[j]
			j--
			shifts++
		}
		arr[j + 1] = key

		fmt.Fprintf(out, "Pass %d: ", i)
		for k := 0; k < n; k++ {
			if k > 0 {
				fmt.Fprintf(out, " ")
			}
			fmt.Fprint(out, arr[k])
		}

		fmt.Fprintf(out, " , ")
		for k := 0; k <= i; k++ {
			if k > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, "%d", arr[k])
		}

		fmt.Fprintf(out, " |")

		for k := i + 1; k < n; k++ {
			fmt.Fprintf(out, " %d", arr[k])
		}

		fmt.Fprintf(out, " , shifts = %d\n", shifts)
	}
}
