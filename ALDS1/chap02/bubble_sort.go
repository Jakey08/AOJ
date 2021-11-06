package main

import "fmt"

//1 bubbleSort(A, N) // N 個の要素を含む 0-オリジンの配列 A
//2   flag = 1 // 逆の隣接要素が存在する
//3   while flag
//4     flag = 0
//5     for j が N-1 から 1 まで
//6       if A[j] < A[j-1]
//7         A[j] と A[j-1] を交換
//8         flag = 1

func bubbleSort(a []int, n int){
	flag := true //逆の隣接要素が存在する
	cnt := 0

	for flag {
		flag = false
		for j := n-1; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				cnt++
				flag = true
			}
		}
	}
	for i := 0; i < n; i++ {
		if i > 0{
			fmt.Printf(" ") //これがないとPresentation errorになる
		}
		fmt.Printf("%d", a[i])
	}
	fmt.Printf("\n")
	fmt.Println(cnt)
}

func main(){
	var length int
	fmt.Scanf("%d", &length)
	A := make([]int, length)
	for i := 0 ; i < length; i++ {
		fmt.Scanf("%d", &A[i])
	}
	bubbleSort(A, length)

}