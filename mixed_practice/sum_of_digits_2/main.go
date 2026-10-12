
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscan(in, &n)

	sum := 0
	for i := 0; i < len(n); i++ {
		sum += int(n[i] - '0')
	}

	fmt.Println(sum)
}
