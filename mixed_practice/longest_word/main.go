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
	s, _ = in.ReadString('\n')

	longest_word, word := 0, 0
	for _, val := range s {
		if val == ' ' || val == '\n' || val == '\r' {
			if word > longest_word {
				longest_word = word
			}
			word = 0
		} else {
			word++
		}
	}
	
	if word > longest_word {
		longest_word = word
	}

	fmt.Fprintln(out, longest_word)
}
