package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var n2 int64
	fmt.Fscan(in, &n2)

	arr2 := make([]int64, n2)
	for i := int64(0); i < n2; i++ {
		fmt.Fscan(in, &arr2[i])
	}

	set1 := make(map[int64]struct{})
	for _, val := range arr1 {
		set1[val] = struct{}{}
	}

	set2 := make(map[int64]struct{})
	for _, val := range arr2 {
		set2[val] = struct{}{}
	}

	intersection_set := make(map[int64]struct{})
	for val := range set1 {
		if _, exists := set2[val]; exists {
			intersection_set[val] = struct{}{}
		}
	}

	fmt.Fprintln(out, len(intersection_set))
	intersection := make([]int64, 0, len(intersection_set))
	for val := range intersection_set {
		intersection = append(intersection, val)
	}
	sort.Slice(intersection, func(i, j int) bool {
		return intersection[i] < intersection[j]
	})
	for _, val := range intersection {
		fmt.Fprintf(out, "%d ", val)
	}
	fmt.Fprintln(out)
}
