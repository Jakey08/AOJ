package main

import "fmt"
func main(){
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	cnt := 0
	for i := a; i <=b; i++ {
		if c % i == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}