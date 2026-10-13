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

	var s, t string
	fmt.Fscan(in, &s, &t)

	ans := make([]byte, 0, 2 * (len(s) + len(t)))
	i, j := 0, 0
	
	for i < len(s) && j < len(t) {
		ans = append(ans, s[i], '-', t[j], '-')
		i++
		j++
	}
	
	for i < len(s) {
		ans = append(ans, s[i], '-')
		i++
	}
	
	for j < len(t) {
		ans = append(ans, t[j], '-')
		j++
	}
	
	if len(ans) > 0 {
		ans = ans[:len(ans)-1]
	}

	fmt.Fprintln(out, string(ans))
}
