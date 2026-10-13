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

	ans := []byte(s)
	for i, val := range s {
		if unicode.IsUpper(val) {
			ans[i] = byte(unicode.ToLower(val)) 
		} else {
			ans[i] = byte(unicode.ToUpper(val))
		}
	}

	fmt.Fprintln(out, string(ans))
}
