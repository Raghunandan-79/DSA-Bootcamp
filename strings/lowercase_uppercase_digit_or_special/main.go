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

	if unicode.IsUpper(ch) {
		fmt.Fprintln(out, "Uppercase")
	} else if unicode.IsLower(ch) {
		fmt.Fprintln(out, "Lowercase")
	} else if unicode.IsDigit(ch) {
		fmt.Fprintln(out, "Digit")
	} else {
		fmt.Fprintln(out, "Special")
	}
}
