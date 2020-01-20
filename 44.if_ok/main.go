package main

import "fmt"

func main() {
	foo := map[string]int{
		"key1": 1,
	}

	if val, ok := foo["key1"]; ok {
		fmt.Printf("key1 exists. Value is %#v\n", val)
	}
	if val, ok := foo["nothing"]; !ok {
		fmt.Printf("Not exists %#v\n", val)
	}
}
