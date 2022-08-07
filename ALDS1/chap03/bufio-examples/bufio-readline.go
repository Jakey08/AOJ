package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio の ReadLineを用いた読み込み方法
// https://qiita.com/tnoda_/items/b503a72eac82862d30c6

var rdr = bufio.NewReaderSize(os.Stdin, 100000)

func readLine() string {
	buf := make([]byte, 0, 100000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func main() {
	x := readLine()
	fmt.Println(x + "\nte\ns\nt\ne\ns")
}
