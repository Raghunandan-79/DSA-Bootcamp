package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	sum := int64(0)
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	fmt.Println(sum)
}
