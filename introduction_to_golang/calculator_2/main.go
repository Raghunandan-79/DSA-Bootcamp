package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int64
	fmt.Fscan(in, &n, &m)

	fmt.Printf("%d + %d = %d\n\n", n, m, (n + m))
	fmt.Printf("%d - %d = %d\n\n", n, m, (n - m))
	fmt.Printf("%d * %d = %d\n\n", n, m, (n * m))
	fmt.Printf("%d / %d = %d\n\n", n, m, (n / m))
	fmt.Printf("%d %% %d = %d\n\n", n, m, (n % m))
}
