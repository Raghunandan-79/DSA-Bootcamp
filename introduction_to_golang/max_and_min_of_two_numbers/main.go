package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a < b {
		fmt.Printf("Min = %d\n", a)
		fmt.Printf("Max = %d\n", b)
	} else {
		fmt.Printf("Min = %d\n", b)
		fmt.Printf("Max = %d\n", a)
	}
}
