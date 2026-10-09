package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x == 0 && y == 0 {
		fmt.Println("Origin")
	} else if y == 0 && x != 0 {
		fmt.Println("x axis")
	} else if x == 0 && y != 0 {
		fmt.Println("y axis")
	} else if x > 0 && y > 0 {
		fmt.Println("1st Quadrant")
	} else if x < 0 && y > 0 {
		fmt.Println("2nd Quadrant")
	} else if x < 0 && y < 0 {
		fmt.Println("3rd Quadrant")
	} else {
		fmt.Println("4th Quadrant")
	}
}
