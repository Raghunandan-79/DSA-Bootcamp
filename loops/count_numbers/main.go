package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	positive, negative, even, odd := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)

		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		}

		if value%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println(positive)
	fmt.Println(negative)
	fmt.Println(even)
	fmt.Println(odd)
}
