package main

import (
	"fmt"
	"strconv"
)

func main() {
	a2 := a.(string)
	i, _ = strconv.Atoi(a2)
	fmt.Println(i) // -> 1234
}

var i int

var a interface{} = "1234"
