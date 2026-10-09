package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	for i := l; i <= r; i++ {
		fmt.Printf("%d ", i)
	}
}
