package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Printf("%d + %d = %d\n", n, m, (n + m))
	fmt.Printf("%d - %d = %d\n", n, m, (n - m))
	fmt.Printf("%d * %d = %d\n", n, m, (n * m))
	fmt.Printf("%d / %d = %d\n", n, m, (n / m))
	fmt.Printf("%d %% %d = %d\n", n, m, (n % m))
}
