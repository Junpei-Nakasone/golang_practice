package main

import "fmt"

func main() {
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int

	p1 = &i
	p1value := p1
	fmt.Println(p1value)

	p2 = &j

	// *p1は数値が入っているので、iは300
	i = *p1 + *p2

	// p2は*intなのでp2 = *p1とすると型の違いでエラーになる(*p1はint)
	p2 = p1

	j = *p2 + i
	fmt.Println(j)
}
