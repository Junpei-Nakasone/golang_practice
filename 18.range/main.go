package main

import "fmt"

func main() {
	l := []string{"ken", "john", "go"}

	for _, v := range l {
		fmt.Println(v)
	}
}
