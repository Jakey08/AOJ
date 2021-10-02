package main

import "fmt"

func main() {
	var (
		S int
		s int
		m int
		h int
	)

	fmt.Scan(&S)

	//1h = 3600s
	//s = 1/3600h
	//1m = 60s
	//s = 1/3600h 1/60

	h = S / 3600
	m = (S % 3600) / 60
	s = S % 60

	fmt.Printf("%d:%d:%d\n", h, m, s) //Printf needs formats
	//https://pkg.go.dev/fmt
}
