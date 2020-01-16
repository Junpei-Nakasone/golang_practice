package main

import "fmt"

type User struct {
	Name string
}

// 値関数で振る舞いを定義
func (u User) Greeting(msg string) {
	fmt.Println("User.Name(関数内):", &(u.Name))
}

// ポインタ関数で振る舞いを定義
func (u *User) Greeting2(msg string) {
	fmt.Println("User.Name(関数内):", &(u.Name))
}

func main() {

	user := User{"Yohei"}

	// 値関数の呼び出し
	fmt.Println("User.Name:", &(user.Name)) // User.Name: 0x1040c128
	user.Greeting("Hello")                  // User.Name(関数内): 0x1040c138（違う！！）

	// ポインタ関数の呼び出し
	fmt.Println("User.Name:", &(user.Name)) // User.Name: 0x1040c128
	user.Greeting2("Hello")                 // User.Name(関数内): 0x1040c128（同じ!!）
}
