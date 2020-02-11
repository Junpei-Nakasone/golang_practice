package main

import "fmt"

type Student struct {
	ID      int
	Profile []profile
}

type profile struct {
	Name string
	Age  int
}

// type team struct {
// 	teamMenber []
// }

func main() {

	student1Profile := profile{
		Name: "st1",
		Age:  10,
	}
	students := []profile{}
	students = append(students, student1Profile)

	student2Profile := profile{
		Name: "st2",
		Age:  5,
	}
	students = append(students, student2Profile)

	fmt.Println(students)

}
