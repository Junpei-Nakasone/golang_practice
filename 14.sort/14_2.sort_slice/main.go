package main

import "sort"

import "fmt"

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"GO", 7},
		{"ALICE", 55},
		{"VERA", 24},
		{"BOB", 75},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("Sorted By Name:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted By Age:", people)
}
