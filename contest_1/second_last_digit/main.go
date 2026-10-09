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

	second_last := n % 10
	n /= 10
	second_last = n % 10
	n /= 10

	fmt.Println(second_last)
}
