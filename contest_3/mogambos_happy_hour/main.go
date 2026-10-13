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

	var time string
	fmt.Fscan(in, &time)

	if time[:2] == time[3:] {
		fmt.Fprintln(out, "Mogambo is happy")
	} else {
		fmt.Fprintln(out, "Mogambo is sad")
	}
}
