package main

import "fmt"

func main() {
	var marks int
	fmt.Scan(&marks)

	if marks > 90 {
		fmt.Println("Excellent")
	} else if marks > 80 {
		fmt.Println("Good")
	} else if marks > 70 {
		fmt.Println("Fair")
	} else if marks > 60 {
		fmt.Println("Meets Expectations")
	} else {
		fmt.Println("Below Par")
	}
}
