package main

import (
	"fmt"
)

func main() {
	for i := 0; ; i++ {
		var x int
		var y int
		fmt.Scan(&x, &y)
		if x == 0 && y == 0 {
			break
		}

		if x > y {
			x, y = y, x
			fmt.Println(x, y)
		} else {
			fmt.Println(x, y)
		}

	}

}