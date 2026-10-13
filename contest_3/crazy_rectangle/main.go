package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	number := 1
	letter := 'a'
	useNumber := true

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if useNumber {
				fmt.Fprint(out, number)
				number = number % 9 + 1
			} else {
				fmt.Fprint(out, string(letter))
				if letter == 'z' {
					letter = 'a'
				} else {
					letter++
				}
			}

			useNumber = !useNumber
		}
		
		fmt.Fprintln(out)
	}
}
