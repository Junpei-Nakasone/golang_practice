package main

import "fmt"

func main() {
	// double型をint型に変換して出力
	a := 4.12

	b := int(a)

	fmt.Println(b)

	m := map[string]int{
		"john": 20,
		"chad": 11,
	}

	fmt.Printf("%T %v\n", m, m)
}
