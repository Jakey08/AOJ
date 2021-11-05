package main

import "fmt"

//Insertion Sort

//Input : The first line of the input includes an integer N,
//the number of elements in the sequence.

//Output : The output consists of N lines.
//Please output the intermediate sequence in a line for each step.
//Elements of the sequence should be separated by single space.


//for i = 1 to A.length-1
//key = A[i]
///* insert A[i] into the sorted sequence A[0,...,j-1] */
//j = i - 1
//while j >= 0 and A[j] > key
//A[j+1] = A[j]
//j--
//A[j+1] = key

 func insertion(a []int){
	 for i, l := 0, len(a); i < l; i++ {
		 if i < l -1 {
			 fmt.Printf("%d ", a[i])
		 }else {
			 fmt.Println(a[i])
		 }
	 }
 }
func main(){
	var n int
	fmt.Scanln(&n)
	a := make([]int, n)
	for i := 0; i< n ; i++ {
		fmt.Scan(&a[i])
	}
	insertion(a)

	for i := 1; i <= len(a)-1 ; i++ {
		temp := a[i]
		j := i -1
		for j >= 0 && a[j] > temp {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = temp
		insertion(a)
	}
}