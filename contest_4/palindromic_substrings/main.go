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
	var s string
	fmt.Fscan(in, &n, &s)

	palindrome := make([][]bool, n)
	for i := range palindrome {
		palindrome[i] = make([]bool, n)
	}

	for length := 1; length <= n; length++ {
		for start := 0; start+length <= n; start++ {
			end := start + length - 1
			palindrome[start][end] = s[start] == s[end] && (length <= 2 || palindrome[start + 1][end - 1])
		}
	}

	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			if palindrome[start][end] {
				fmt.Fprintln(out, s[start:end + 1])
			}
		}
	}
}
