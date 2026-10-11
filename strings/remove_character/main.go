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

	var s string
	fmt.Fscan(in, &s)
	var x string
	fmt.Fscan(in, &x)
	ch := rune(x[0])

	result := make([]rune, 0, len([]rune(s)))
	for _, val := range s {
		if val != ch {
			result = append(result, val)
		}
	}

	fmt.Fprintln(out, string(result))
}
