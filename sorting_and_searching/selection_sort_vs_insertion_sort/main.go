package main

import (
	"bufio"
	"fmt"
	"os"
)

func countInsertionShifts(a []int) int64 {
	var shifts int64
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			shifts++
			j--
		}
		a[j+1] = key
	}
	return shifts
}

func countSelectionSwaps(a []int) int64 {
	var swaps int64
	for i := 0; i < len(a)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			a[i], a[minIndex] = a[minIndex], a[i]
			swaps++
		}
	}
	return swaps
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for t > 0 {
		var n int
		fmt.Fscan(in, &n)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		insertionShifts := countInsertionShifts(append([]int(nil), arr...))
		selectionSwaps := countSelectionSwaps(append([]int(nil), arr...))

		if insertionShifts < selectionSwaps {
			fmt.Fprintln(out, "Insertion Sort")
		} else if selectionSwaps < insertionShifts {
			fmt.Fprintln(out, "Selection Sort")
		} else {
			fmt.Fprintln(out, "Tie")
		}

		t--
	}
}
