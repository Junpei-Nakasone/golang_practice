package main

import "fmt"

type repository interface {
	Test(int) int
}

func Test(x int) int {
	x = x + 1
	return x
}

func (r repository) Res() {
	r.Test(1)
}

func main() {
	a := Res()
	fmt.Println(a)
}
