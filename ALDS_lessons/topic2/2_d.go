package main

import (
	"fmt"
)

func main() {
	var (
		W int
		H int
		x int
		y int
		r int
	)

	fmt.Scan(&W, &H, &x, &y, &r)
	if 0 <= x - r && x - r < W && x + r <= W && 0 <= y - r && y -r < H && y + r <= H {
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}
