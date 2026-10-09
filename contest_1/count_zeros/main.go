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

	if n == 0 {
		fmt.Println(1)
		return
	}

	cnt := 0
	for n != 0 {
		last_digit := n % 10
		if last_digit == 0 {
			cnt++
		}
		n /= 10
	}
	fmt.Println(cnt)
}
