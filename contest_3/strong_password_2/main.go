package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	s, _ := in.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	s = strings.TrimSuffix(s, "\r")

	lower, upper, digit, special := false, false, false, false
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			lower = true
		case r >= 'A' && r <= 'Z':
			upper = true
		case r >= '0' && r <= '9':
			digit = true
		default:
			special = true
		}
	}

	categories := 0
	for _, present := range []bool{lower, upper, digit, special} {
		if present {
			categories++
		}
	}

	switch {
	case categories == 4:
		fmt.Fprintln(out, "Strong")
	case categories >= 3:
		fmt.Fprintln(out, "Moderate")
	default:
		fmt.Fprintln(out, "Weak")
	}
}
