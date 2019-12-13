package main

import "fmt"

func main() {
	os := "windows"
	switch os {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("Default")
	}
}
