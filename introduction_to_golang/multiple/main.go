package main

import "fmt"

func main() {
	var n, m int64
	fmt.Scan(&n, &m)

	if m % n == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
