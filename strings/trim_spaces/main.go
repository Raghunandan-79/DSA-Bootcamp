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

	for _, val := range s {
		if val != ' ' {
			fmt.Fprintf(out, "%c", val)
		}
	}
}
