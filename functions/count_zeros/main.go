package main

import "fmt"

func count_zeros(n int64) int64 {
	if n == 0 {
		return 1
	}

	cnt := int64(0)
	
	for n != 0 {
		last_digit := n % 10
		
		if last_digit == 0 {
			cnt++
		}

		n /= 10
	}

	return cnt
}

func main() {
	var n int64
	fmt.Scan(&n)

	fmt.Println(count_zeros(n))	
}
