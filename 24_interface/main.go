package main

import "fmt"

// interfaceは実装するコードは書かずにメソッドの型のみ宣言
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

// 構造体PersonにSay()を実装
// 関数内でPersonの値を書き換えるので、*で指定
// 返す値にstringを指定しているのは、関数内にreturnがなければ
// 必要ない。
func (p *Person) Say() string {

	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

// 引数でinterfaceを受け取る関数
func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("RUN")
	} else {
		fmt.Println("GETOUT")
	}

}

func main() {
	// 変数mikeにinterfaceのHumanを格納し、
	// Humanに文字列Mikeを持つ構造体Personを入れる
	// 値が書き換わったPersonを受け取るため&で指定
	var mike Human = &Person{"Mike"}
	// Humanという構造体にPersonを入れたらSay()を持ってないといけない
	DriveCar(mike)

	var x Human = &Person{"X"}

	DriveCar(x)
}
