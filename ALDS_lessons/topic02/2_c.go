package main

import "fmt"

func main() {
	var (
		a   int
		b   int
		c   int
		tmp int
	)

	fmt.Scan(&a, &b, &c)
	if a > b {
		tmp = a
		a = b
		b = tmp
	}
	if b > c {
		tmp = b
		b = c
		c = tmp
	}
	if a > b {
		tmp = a
		a = b
		b = tmp
	}
	fmt.Println(a, b, c)
}
