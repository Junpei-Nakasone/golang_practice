package main

import (
	"fmt"
)

func ReverseInt(n int) int {
	// 出力する数値の変数を宣言
	new_int := 0

	// 引数が0より大きい間のforループを作成
	for n > 0 {
		// 引数を10で割った余りを変数reminderに格納
		reminder := n % 10
		// 出力用の変数を10倍にする
		new_int *= 10
		// 出力用の変数にreminderを足す
		new_int += reminder
		// 引数を10で割る
		n /= 10
	}
	// 出力用の変数をreturnする
	return new_int
}

func main() {
	n := ReverseInt(123)
	fmt.Println(n)

}
