package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)

	compressed := make([]byte, 0, len(s))
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}

		count := j - i
		compressed = append(compressed, s[i])
		if count > 1 {
			compressed = append(compressed, byte('0'+count))
		}
		i = j
	}

	fmt.Fprintln(out, string(compressed))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for t > 0 {
		solve(in, out)
		t--
	}
}
