package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > b && a > c {
		if b > c {
			fmt.Printf("Min = %d\n", c)
			fmt.Printf("Max = %d\n", a)
		} else {
			fmt.Printf("Min = %d\n", b)
			fmt.Printf("Max = %d\n", a)
		}
	} else if b > a && b > c {
		if a > c {
			fmt.Printf("Min = %d\n", c)
			fmt.Printf("Max = %d\n", b)
		} else {
			fmt.Printf("Min = %d\n", a)
			fmt.Printf("Max = %d\n", b)
		}
	} else {
		if a > b {
			fmt.Printf("Min = %d\n", b)
			fmt.Printf("Max = %d\n", c)
		} else {
			fmt.Printf("Min = %d\n", a)
			fmt.Printf("Max = %d\n", c)
		}
	}
}
