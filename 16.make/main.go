package main

import "fmt"

func main() {
	// makeで値が入ってないスライスをlengthとcapacityを定義して宣言できる

	// makeの第一引数はデータ型、第二引数はlength, 第３引数はcapacity
	n := make([]int, 3, 5)

	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	c := make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
