package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin) // Scannerを用意する

func nextInt() int {
	sc.Scan()                       // 入力を読み込む
	i, e := strconv.Atoi(sc.Text()) //入力された文字列を整数に変換
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords) //ScanWordsでspaceを削除する
	n := nextInt()
	x := 0
	for i := 0; i < n; i++ {
		x += nextInt()
	}
	fmt.Println(x)
}
