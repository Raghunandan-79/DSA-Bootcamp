package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	if n % 100 == 0 {
		if n % 400 == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		if n % 4 == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
