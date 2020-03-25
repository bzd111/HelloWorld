package main

import (
	"fmt"
	"time"
)

func order() {
	// 后入先出，
	// 因为后面的defer,会追加到链表的最前面
	for i := 0; i < 5; i++ {
		defer fmt.Println("i: ", i)
	}
}

func arguments() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func since() {
	startAt := time.Now()
	// defer调用是就会拷贝变量
	defer fmt.Println("direct print", time.Since(startAt))
	defer func() {
		fmt.Println("in func: ", time.Since(startAt))
	}()
	time.Sleep(time.Second)
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	order()
	since()
	arguments()
	fmt.Println(c())
}
