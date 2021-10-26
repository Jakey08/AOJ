package main

import "fmt"

func main(){

	for i:= 1; ; i++ {
		var x int
		fmt.Scan(&x)

		if x == 0 {
			break
		}else{
			fmt.Printf("Case %d: %d\n" ,i,x) //\n is needed cause it's doesn't start new line.
		}
	}
}