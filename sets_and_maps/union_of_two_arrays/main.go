package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n1 int64
	fmt.Fscan(in, &n1)

	arr1 := make([]int64, n1)
	for i := int64(0); i < n1; i++ {
		fmt.Fscan(in, &arr1[i])
	}
	slices.Sort(arr1)

	var n2 int64
	fmt.Fscan(in, &n2)

	arr2 := make([]int64, n2)
	for i := int64(0); i < n2; i++ {
		fmt.Fscan(in, &arr2[i])
	}
	slices.Sort(arr2)

	union := []int64{}
	i, j := int64(0), int64(0)
	for i < n1 && j < n2 {
		if arr1[i] <= arr2[j] {
			if len(union) == 0 || union[len(union)-1] != arr1[i] {
				union = append(union, arr1[i])
			}
			i++
		} else {
			if len(union) == 0 || union[len(union)-1] != arr2[j] {
				union = append(union, arr2[j])
			}
			j++
		}
	}

	for i < n1 {
		if len(union) == 0 || union[len(union)-1] != arr1[i] {
			union = append(union, arr1[i])
		}
		i++
	}

	for j < n2 {
		if len(union) == 0 || union[len(union)-1] != arr2[j] {
			union = append(union, arr2[j])
		}
		j++
	}

	fmt.Fprintln(out, len(union))
	for i, value := range union {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, value)
	}
	fmt.Fprintln(out)
}
