package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	minimum := arr[0]
	minimum_idx := 0
	for i := 1; i < n; i++ {
		if arr[i] < minimum {
			minimum = arr[i]
			minimum_idx = i
		}
	}

	fmt.Println(minimum, minimum_idx + 1)
}
