package main

import "fmt"

// 配列は固定長だが、スライスは可変長で柔軟。
//

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 12}

	var s []int = primes[0:1]
	fmt.Println(s)
}
