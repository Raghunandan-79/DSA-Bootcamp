package main

import "fmt"

func main() {
	var n, m int64
	fmt.Scan(&n, &m)

	fmt.Printf("%d", (n % 10) + (m % 10))
}
