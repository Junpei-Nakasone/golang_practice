package main

import "fmt"

func main() {
	var a int
	a = 1
	one(a)
	fmt.Println("a:", a)
	fmt.Println("&a:", &a)
	var b *int
	b = &a
	fmt.Println("*b:", *b)
	two(b)
	fmt.Println("b:",b)
	fmt.Println("*b", *b)
	fmt.Println("a:",a)
}

func one(x int) {
	x = 2
}

func two(x *int) {
	*x = 2
}