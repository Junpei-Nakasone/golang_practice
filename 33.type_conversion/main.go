package main

// 文字列からint, floatへは直接変換できないのでstrconvパッケージを使う必要がある

import "fmt"

import "strconv"

func main() {
	var s string = "14"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("%T %v", i, i)
}
