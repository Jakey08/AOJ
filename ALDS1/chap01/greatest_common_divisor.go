package main

import "fmt"

func gcd(x int, y int )  int{
	if x < y {
		swap(x,y)
	}
	for y > 0 {
		r := x % y
		x = y
		y = r
	}
	return x
}

func swap(a int, b int){

	b, a = a, b
}


func main() {
	var (
		x int
		y int
	)
	fmt.Scanln(&x, &y)
	  t := gcd(x, y)
	fmt.Println(t)
}