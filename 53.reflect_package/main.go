package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name  string
	Age   int
	Which bool
}

func main() {
	a := Test{
		"MenA",
		10,
		true,
	}

	a2 := reflect.ValueOf(a)
	a3 := a2.Type()
	fmt.Println(a2)
	fmt.Println(a3)

}
