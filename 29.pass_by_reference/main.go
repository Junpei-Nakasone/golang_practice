package main

import "fmt"

type ST1 struct {
	f1 string
	f2 map[string]string
}

func (s ST1) a() {

}

type IF1 interface {
	a()
}

func F1(s ST1) {
	s.f1 = "f1"
	s.f2["f2"] = "test"
}

func main() {
	fmt.Println("*** 1.passByReference")
	s1 := ST1{"", map[string]string{}}
	fmt.Printf("1: s1 = %#v\n", s1)
	F1(s1)
	fmt.Printf("2: s1 = %#v\n", s1)

}
