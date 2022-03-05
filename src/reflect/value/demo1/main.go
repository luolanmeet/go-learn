package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) String() {
	println("User:", user.Id, user.Age, user.Name)
}

func Info(o interface{}) {
	// 获取 Value 信息
	v := reflect.ValueOf(o)
	// 通过 Value 获取 Type
	t := v.Type()
	// 类型名称
	println(t.Name()) // User

	// 访问接口字段名、字段类型和字段值信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		// 类型查询
		switch val := value.(type) {
		case int:
			fmt.Printf(" %6s: %v=%d\n", field.Name, field.Type, val)
		case string:
			fmt.Printf(" %6s: %v=%s\n", field.Name, field.Type, val)
		default:
			fmt.Printf(" %6s: %v=%s\n", field.Name, field.Type, val)
		}
	}
}

func main() {
	u := User{1, "cck", 27}
	Info(u)
}
