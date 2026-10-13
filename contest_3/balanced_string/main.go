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

	cnt_a, cnt_b := 0, 0
	for _, val := range s {
		if val == 'a' {
			cnt_a++
		} else {
			cnt_b++
		}
	}

	if cnt_a == cnt_b {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
