package convert

import (
	tdomain "ZO_School_Wares/common/time/domain"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9]([A-Z])")

// StructToMap は構造体からMapに変換します(ただしゼロ値(nil)の場合はmapに含みません)
func StructToMap(v interface{}) map[string]interface{} {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Struct:
		m := make(map[string]interface{})
		// 構造体のフィールド全てをmapに格納する
		setValueToMap(m, v)
		return m
	default:
		// 構造体以外はnilを返す
		return nil
	}
}

// setValueToMap　は構造体の値をmapに格納します
func setValueToMap(m map[string]interface{}, v interface{}) {
	rv := reflect.ValueOf(v)
	rt := rv.Type()
	// 構造体のフィールドごとの値をmapに格納する
	for i := 0; i < rv.NumField(); i++ {
		fName := rt.Field(i).Name
		fValue := rv.FieldByName(fName)
		var value interface{}
		fbnt := fValue.Type()

		// 種類別に値の取得方法を切り替える
		switch reflect.TypeOf(fValue.Interface()).Kind() {
		case reflect.Ptr:
			// ポインタ型の場合
			if fValue.IsNil() {
				// ゼロ値(nil)の場合は以降の処理をスキップする
				continue
			}
			value = reflect.Indirect(fValue).Interface()
		case reflect.Struct:
			value = reflect.Indirect(fValue).Interface()
			switch fbnt {
			case reflect.TypeOf(time.Time{}),
				reflect.TypeOf(tdomain.Date{}),
				reflect.TypeOf(tdomain.Time{}):
				// time.Time型のば良いはdefaultと同様の処理を行う
				if isValueZero(value, fbnt) {
					// ポインタ型以外かゼロ値の場合は以降の処理をスキップする
					continue
				}
			default:
				// その他の型の場合
				value = reflect.Indirect(fValue).Interface()
				if isValueZero(value, fbnt) {
					// ポインタ型以外かつゼロ値の場合は以降の処理をスキップする
					continue
				}
			}
			// keyはスネークケースに変換する
			key := ToSnakeCase(fName)
			m[key] = value
		}
	}
}

// ToSnakeCaseはキャメルケースをスネークケースに変換する
func ToSnakeCase(s string) string {
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func isValueZero(v interface{}, t reflect.Type) bool {
	return v == reflect.Zero(t).Interface()
}
