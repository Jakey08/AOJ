package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := float64(a) / float64(b)

	fmt.Printf("%d %d %f\n", a/b, a%b, c)
}
