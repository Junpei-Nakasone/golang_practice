package main

import (
	"reflect"
)

func main() {
	type StoMTest struct {
		Name string
		Age  int
	}

	a := StoMTest{
		"NoPointer",
		1,
	}

	b := structToMap(a)

}

func structToMap(v interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	setValueToMap(m, v)
	return m
}

func setValueToMap(m map[string]interface{}, v interface{}) {
	rv := reflect.ValueOf(v)
	rt := rv.Type()
}
