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

	length_of_string := len(s)
	lowercase_count := 0
	uppercase_count := 0
	digit_count := 0
	special_count := 0

	for _, val := range(s) {
		if unicode.IsLower(val) {
			lowercase_count++
		} else if unicode.IsUpper(val) {
			uppercase_count++
		} else if unicode.IsDigit(val) {
			digit_count++
		} else {
			special_count++
		}
	}

	if length_of_string == 10 && lowercase_count >= 1 && uppercase_count >= 1 && digit_count >= 1 && special_count >= 1 {
		fmt.Fprintln(out, "Strong")
	} else {
		fmt.Fprintln(out, "Weak")
	}
}
