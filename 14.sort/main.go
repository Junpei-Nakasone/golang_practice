package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	// 宣言された並びのまま表示される
	fmt.Println(i, s, p) //[5 3 2 7] [d a f] [{Nancy 20} {Vera 40} {Mike 30} {Bob 50}]

	// sortパッケージのInts関数で数値をソートする
	sort.Ints(i)
	fmt.Println(i) // [2 3 5 7]

	// sortパッケージのStrings関数で文字列をソートする
	sort.Strings(s)
	fmt.Println(s) //[a d f]

	// sortパッケージのSlice関数でsliceをソートする
	// iとjで二つfor roopを回してるような形になり、Nameの昇順にソートされる
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	fmt.Println(p) //[{Bob 50} {Mike 30} {Nancy 20} {Vera 40}]
}
