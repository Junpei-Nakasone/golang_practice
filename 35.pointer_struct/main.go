package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}
type Person2 struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("OK")
	} else {
		fmt.Println("NO")
	}
}

func main() {

	var mike Human = &Person{Name: "Mike"}
	// mike.Say()
	DriveCar(mike)
}
