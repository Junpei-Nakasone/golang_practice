package main

import (
	"fmt"
)

func main() {

	// 変数iに100を格納
	var a int = 100
	fmt.Println("aのアドレス", &a)

	// 変数bに200を格納
	var b int = 200

	// ポインタ型の変数p1を宣言
	var p1 *int
	// p1のアドレス

	// ポインタ型の変数p2を宣言
	var p2 *int

	// p1に変数aのアドレスを格納
	p1 = &a

	fmt.Println("p1:aのアドレス", p1)
	fmt.Println("*p1:aの値", *p1)
	fmt.Println("&p1:p1のアドレス", &p1)

	// p2に変数bのアドレスを格納
	p2 = &b

	fmt.Println("p2:bのアドレス", p2)
	fmt.Println("*p2:bの値", *p2)
	fmt.Println("&p2:p2のアドレス", &p2)

	// 変数aに*p1でaのもつ値(100)、*p2でbももつ値(200)を格納
	a = *p1 + *p2
	fmt.Println("*p1と*p2を足したa:", a)

	fmt.Println("*p1: *p1はaの値なので300に変わってる", *p1)

	p2 = p1
	fmt.Println("p2: p1の値(aのアドレス)を格納", p2)
	fmt.Println("*p2: ", *p2)

	b = *p2 + a

	fmt.Println(b)

}
