package main

import "fmt"

// カスタムエラーを作る時にはポインタで&を渡す
// Go組み込みのエラーと違うアドレスを指すと挙動が変わるため

type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.UserName)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{UserName: "mike"}

}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
