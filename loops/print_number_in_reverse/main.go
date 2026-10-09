package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}

	for n > 0 {
		fmt.Printf("%d", n % 10)
		n /= 10
	}
}
