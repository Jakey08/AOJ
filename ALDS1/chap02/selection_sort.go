package main

import "fmt"

func selectionSort(a []int, n int) {
	var cnt int
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if a[j] < a[minj] {
				minj = j
			}
		}
		if i != minj {
			a[i], a[minj] = a[minj], a[i]
			cnt++
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", a[i])
	}
	fmt.Printf("\n")
	fmt.Println(cnt)
}

func main(){
	var x int
	fmt.Scanf("%d", &x) //&つけ忘れると、格納できない

	a := make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Scanf("%d",&a[i])
	}
	selectionSort(a, x)
}