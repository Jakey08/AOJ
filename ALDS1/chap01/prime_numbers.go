package main

//Input : The first line contains an integer N,
//the number of elements in the list.
//N numbers are given in the following lines.

//Output
//Print the number of prime numbers in the given list.

// How to judge the prime numbers?
// 25 = 2 * 12 + 1
// 12 = 2 * 6 + 0

// 33 = 2 + 16 + 1
// 79 = 2 * 39 + 1
// 39 = 2 * 19 + 1
// 19 = 2 * 9 + 1
// 50 = 2 * 5^2
// 101 = 2 * 2 + 5^2 + 1
// x := n % 2
// if x == 1 braek, this is prime number
//else that's not

//Method : p <= square(x)
//isprime(x) if x == 2 return true
// if x < 2 or x % 2 == 0 return false
// i = 3 for i <= square(x)
// if x % i == 0 return false
// i = i + 2 return true

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {

	if x == 2{
		return true
	}else if x < 2 || x % 2 == 0 {
		return false
	}
	n := int(math.Sqrt(float64(x)))
	for i :=3; i <= n; i += 2 {
		if x % i == 0 {
			return false
		}
	}
	return true
}

func main(){
	var (
		n int
		x int
		cnt int
	)
	fmt.Scan(&n)
	for i := 0; i < n ; i++ {
		fmt.Scan(&x)
		if isPrime(x){
			cnt += 1
		}
	}
	fmt.Println(cnt)
}
