package main

import (
	"fmt"
)

func main() {
	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	//fmt.Println("%d %d", a*b, 2*(a+b)) %d %dãããªã
	fmt.Println(a*b, 2*(a+b))

}
