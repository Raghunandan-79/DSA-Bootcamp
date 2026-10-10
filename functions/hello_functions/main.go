package main

import (
	"bufio"
	"fmt"
	"os"
)

func say_hello(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("I am learning functions")
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	say_hello(n)
}
