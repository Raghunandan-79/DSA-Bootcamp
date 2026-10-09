package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", 5, i, (5 * i))
	}
}
