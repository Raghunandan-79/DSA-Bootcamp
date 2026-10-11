package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	result := make([]rune, len(s))
	for i, val := range s {
		if unicode.IsUpper(val) {
			result[i] = unicode.ToLower(val)
		} else {
			result[i] = unicode.ToUpper(val)
		}
	}

	fmt.Fprintln(out, string(result))
}
