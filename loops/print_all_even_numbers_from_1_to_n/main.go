package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
