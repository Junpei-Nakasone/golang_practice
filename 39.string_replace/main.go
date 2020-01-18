package main

import "fmt"

import "strings"

func main() {
	str := "Go Java JS Typescript CoffeeScript"

	fmt.Printf("str : %s\n", str)

	replaced1 := strings.Replace(str, "Java", "Replaced", 1)
	fmt.Printf("replaced1 : %s\n", replaced1)
}
