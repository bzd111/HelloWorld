package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 10
	typeOf := reflect.TypeOf(num)
	value := reflect.ValueOf(num)
	fmt.Println("typeOf: ", typeOf)
	fmt.Println("value: ", value)

	// 修改值的变量
	// 获取指针
	value = reflect.ValueOf(&num)
	// 获取指针指向的变量
	value.Elem().SetInt(20)
	fmt.Println("value: ", num)

}
