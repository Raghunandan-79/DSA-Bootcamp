package main

import "fmt"

func main() {
	var ch rune
	fmt.Scanf("%c", &ch)

	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
