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

	var first_name_1, last_name_1 string
	fmt.Fscan(in, &first_name_1, &last_name_1)

	var first_name_2, last_name_2 string
	fmt.Fscan(in, &first_name_2, &last_name_2)

	if last_name_1 == last_name_2 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
