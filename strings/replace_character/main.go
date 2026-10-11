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

	var s string
	fmt.Fscan(in, &s)
	var ch1, ch2 string
	fmt.Fscan(in, &ch1, &ch2)

	s = strings.ReplaceAll(s, ch1, ch2)
	fmt.Fprintln(out, s)
}
