package main

import "fmt"

func main() {
	test := []int{1, 2, 3}
	a := append(test, 4)
	fmt.Println(a)

	test2 := []int{5, 6, 7}

	test = append(test, test2...)

	fmt.Println(test)
}
