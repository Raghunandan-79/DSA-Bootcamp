package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var length, breadth int64
	fmt.Fscan(in, &length, &breadth)
	fmt.Println("Area =", (length * breadth))
	fmt.Println("Perimeter =", (2 * (length + breadth)))
}
