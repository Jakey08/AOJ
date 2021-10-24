package main

//fmt.scan is easy to code, but the process isn't fast compared to 'bufio'.
//cf.) https://qiita.com/tnoda_/items/b503a72eac82862d30c6
// 1. the most easiest way → fmt.scan
// 2. If you want to read a bunch of variables, use bufio Scanner
// 3. If you want to read long variables like >64*10^3, use bufio ReadLine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil{
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	  W := readInt()
	  H := readInt()
	  x := readInt()
	  y := readInt()
	  r := readInt()

	if 0 <= x - r && x - r < W && x + r <= W && 0 <= y - r && y -r < H && y + r <= H {
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}
