package main

import "fmt"

func main() {
	// 最小値を探すコード
	l := []int{200, 500, 34, 5}

	var minimum int
	for k, v := range l {
		if k == 0 {
			minimum = v
		}
		if minimum >= v {
			minimum = v
		}
	}
	fmt.Println(minimum)

	// 合計値を出力するコード
	m := map[string]int{
		"apple":  100,
		"banana": 800,
		"orange": 200,
	}

	var total int
	for _, v := range m {
		total += v
	}
	fmt.Println(total)
}
