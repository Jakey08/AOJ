package main

import (
	"fmt"
)

func main() {
	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	//fmt.Println("%d %d", a*b, 2*(a+b)) %d %dいらない
	fmt.Println(a*b, 2*(a+b))

}
