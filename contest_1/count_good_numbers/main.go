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
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if arr[i] == 0 || 18 % arr[i] == 0 || arr[i] % 45 == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
