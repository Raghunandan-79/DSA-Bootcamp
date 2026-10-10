package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sum := int64(0)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
