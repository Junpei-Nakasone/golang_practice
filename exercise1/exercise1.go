package main

import "fmt"

// type Student struct {
// 	Name string
// 	Data []data
// }

type WhereList struct {
	column   string
	operator string
	value    string
}

// data1 := data{
// 	study:  "good",
// 	sports: "good",
// }

// 整形処理
func test(column, operator, value string) string {
	return column + " " + operator + " " + value
}

func main() {
	res := test("staff_id", "=", "1")
	fmt.Println(res)
	// var data2 []data
	// data2 = append(data2, data1)

	// data3 := data{
	// 	study:  "nice",
	// 	sports: "nice",
	// }

	// data2 = append(data2, data3)

	// user := Student{
	// 	Name: "User1",
	// 	Data: data2,
	// }

	// fmt.Println(user)
}
