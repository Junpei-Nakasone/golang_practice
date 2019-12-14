package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T, %v\n", i, i)
}
