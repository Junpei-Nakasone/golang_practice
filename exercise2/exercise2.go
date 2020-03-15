package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func p(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	type StoMTest struct {
		Name string
		Age  *int
	}
	i := 2
	p := &i

	a := StoMTest{
		"NoPointer",
		p,
	}

	b := structToMap(a)
	fmt.Println(b)

}

func structToMap(v interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	setValueToMap(m, v)
	return m
}

func setValueToMap(r map[string]interface{}, v interface{}) {
	// 構造体のValueOfをrvに格納
	rv := reflect.ValueOf(v)

	// 構造体の中身のTypeをrtに格納
	rt := rv.Type()

	// 構造体の中身の数だけforループを回す
	for i := 0; i < rv.NumField(); i++ {

		// Typeのi番目のFieldのNameをfNameに格納
		fName := rt.Field(i).Name

		// 構造体のValueOfのfNameのFieldByNameをfValueに格納
		fValue := rv.FieldByName(fName)

		// fValueを引数にreflect.Indirectを呼び、.Interface()でどの型でもvalue変数に格納する
		value := reflect.Indirect(fValue).Interface()

		// fNameをスネークケースにして変数keyに格納
		key := ToSnakeCase(fName)

		// mapの[key]とvalueを指定
		r[key] = value

	}
}

// ToSnakeCase makes snakeCase
func ToSnakeCase(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
