package main

import (
	"fmt"
	"strconv"
)

func main() {
	a2 := a.(string)
	i, _ = strconv.Atoi(a2)
	fmt.Println(i) // -> 123
}

var i int

// var s string = "123"

var a interface{} = "1234"
