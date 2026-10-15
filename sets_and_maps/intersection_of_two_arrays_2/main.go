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

	var n1 int
	fmt.Fscan(in, &n1)

	counts := make(map[int64]int)
	for i := 0; i < n1; i++ {
		var value int64
		fmt.Fscan(in, &value)
		counts[value]++
	}

	var n2 int
	fmt.Fscan(in, &n2)
	intersection := make([]int64, 0)
	for i := 0; i < n2; i++ {
		var value int64
		fmt.Fscan(in, &value)
		if counts[value] > 0 {
			intersection = append(intersection, value)
			counts[value]--
		}
	}

	sort.Slice(intersection, func(i, j int) bool {
		return intersection[i] < intersection[j]
	})

	fmt.Fprintln(out, len(intersection))
	for _, value := range intersection {
		fmt.Fprint(out, value, " ")
	}
	fmt.Fprintln(out)
}
