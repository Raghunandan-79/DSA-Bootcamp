package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	arr1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr1[i])
	}

	var m int
	fmt.Fscan(in, &m)

	arr2 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &arr2[i])
	}

	sort.Ints(arr2)

	first := true
	for k := 0; k < n; k++ {
		val := arr1[k]
		idx := sort.SearchInts(arr2, val)

		if idx < len(arr2) && arr2[idx] == val {
			if !first {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, val)
			first = false

			arr2 = append(arr2[:idx], arr2[idx+1:]...)
		}
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
		solve(in, out)
		t--
	}
}
