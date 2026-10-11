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

	var ch rune
	fmt.Fscanf(in, "%c", &ch)

	fmt.Fprintf(out, "%d", ch)
}
