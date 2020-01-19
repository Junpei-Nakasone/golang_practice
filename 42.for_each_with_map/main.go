package main

import "fmt"

func main() {
	myList := map[string]string{
		"dog":      "woof",
		"cat":      "meow",
		"hedgehog": "sniff",
	}

	for animal, noise := range myList {
		fmt.Println("The", animal, "went", noise)
	}
}
