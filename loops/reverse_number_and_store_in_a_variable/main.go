package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}
	
	ans := 0
	for n > 0 {
		last_digit := n % 10
		ans = (ans * 10) + last_digit
		n /= 10
	}
	fmt.Println(ans)
}
