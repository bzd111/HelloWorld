package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}

	for i, _ := range arr {
		newArr = append(newArr, &arr[i])
	}
	for _, v := range arr {
		// 使用中间变量
		_v := v
		newArr = append(newArr, &_v)
	}
	// 错误的写法，newArr保存都是最后一个值
	// for _, v := range arr {
	//     newArr = append(newArr, &v)
	// }
	for _, v := range newArr {
		fmt.Println(*v)
	}
}
