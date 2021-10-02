package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a == b")
	}
}
