package main

import "fmt"

func main() {
	var n, f int64
	fmt.Scan(&n, &f)

	if n % f == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
