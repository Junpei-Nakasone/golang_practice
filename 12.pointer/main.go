package main

import "fmt"

func main() {
	// 変数nに数値100を格納
	var n int
	n = 100

	// 変数nを出力
	fmt.Println(n) // 100が出力される

	// &を先につけることで変数があるメモリのアドレスを指定する
	fmt.Println(&n) // メモリのアドレスが表示される

	// 変数pに*でintのポインタ型にして、&nを格納する
	var p *int
	p = &n

	// メモリのアドレスが表示される
	fmt.Println(p)

	// *pにすることで中身を参照する
	fmt.Println(*p) // 100が表示される

	// &をつけて、nのアドレスを渡し、one関数は*がついてるので、アドレスの中身を書き換える
	one(&n)
	// 1が出力される
	fmt.Println(n)
	fmt.Println(&n)

	noOne(n)

	fmt.Println(n)
	fmt.Println(&n)
}

// *をつけて中身を参照することで、xの値が１になる
func one(x *int) {
	// xの実態の中に1を格納する。dereferenceと呼ばれる
	*x = 1
	fmt.Println(&x)
}

// *をつけずにxの値を変更しても、中身に影響はない
func noOne(x int) {
	x = 2
	fmt.Println(&x)
	fmt.Println(x)
}
