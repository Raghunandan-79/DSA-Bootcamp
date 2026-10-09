package main

import "fmt"

func main() {
	var marks int
	fmt.Scan(&marks)

	if marks >= 35 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
