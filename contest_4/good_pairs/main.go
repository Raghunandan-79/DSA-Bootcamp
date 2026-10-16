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

	var n int
	var k int64
	fmt.Fscan(in, &n, &k)

	freq := make(map[int64]int64)
	var answer int64
	var seen int64
	for i := 0; i < n; i++ {
		var value int64
		fmt.Fscan(in, &value)
		
		if k == 0 {
			if value == 0 {
				answer += seen
			}
		} else if value%k == 0 {
			answer += freq[value/k]
		}

		freq[value]++
		seen++
	}

	fmt.Fprintln(out, answer)
}
