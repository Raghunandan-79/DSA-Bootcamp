package main

import "fmt"

func hcf(a int64, b int64) int64 {
	for a > 0 && b > 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}

	if b == 0 {
		return a
	}

	return b
}

func main() {
	var a, b int64
	fmt.Scan(&a, &b)

	fmt.Println(hcf(a, b))
}
