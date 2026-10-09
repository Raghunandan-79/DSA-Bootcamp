package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	if n % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
