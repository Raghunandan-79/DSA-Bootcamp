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

	var n, m int
	fmt.Fscan(in, &n, &m)

	arr1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr1[i])
	}

	arr2 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &arr2[i])
	}
	
	result := make([]int, 0, n + m)
	i, j := 0, 0
	for i < n && j < m {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < n {
		result = append(result, arr1[i])
		i++
	}

	for j < m {
		result = append(result, arr2[j])
		j++
	}

	for i := 0; i < len(result); i++ {
		fmt.Fprintf(out, "%d ", result[i])
	}
}
