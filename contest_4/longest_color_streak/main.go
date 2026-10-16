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

	var n int64
	fmt.Fscan(in, &n)

	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	longestStreak, streak := int64(1), int64(1)
	for i := int64(1); i < n; i++ {
		if arr[i - 1] == arr[i] {
			streak++
			if streak > longestStreak {
				longestStreak = streak
			}
		} else {
			streak = 1
		}
	}

	fmt.Fprintln(out, longestStreak)
}
