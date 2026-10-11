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

	var ch rune
	fmt.Fscanf(in, "%c", &ch)

	ch = unicode.ToLower(ch)

	fmt.Fprintf(out, "%c", ch)
}
