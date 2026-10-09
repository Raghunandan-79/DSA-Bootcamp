package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println(fact)
}
