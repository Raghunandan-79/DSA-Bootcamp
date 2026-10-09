package main

import "fmt"

func reverse(n int64) int64 {
	ans := int64(0)
	
	for n > 0 {
		last_digit := n % 10
		ans = (ans * 10) + last_digit
		n /= 10
	}

	return ans
}

func main() {
	var n int64
	fmt.Scan(&n)

	if reverse(n) == n {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
