package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	maximum_element := arr[0]
	maximum_element_idx := 0
	for i := 0; i < n; i++ {
		if arr[i] > maximum_element {
			maximum_element = arr[i]
			maximum_element_idx = i
		}
	}

	fmt.Println(maximum_element, maximum_element_idx+1)
}
