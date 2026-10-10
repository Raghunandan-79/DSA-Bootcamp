package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	arr := make([]int64, n)

	for i := int64(0); i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}
