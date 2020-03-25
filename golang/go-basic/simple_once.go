package main

import (
	"fmt"
	"sync"
)

func once() {
	// sync.Once.Do 方法中传入的函数只会被执行一次，哪怕函数中发生了 panic；
	// 两次调用 sync.Once.Do 方法传入不同的函数也只会执行第一次调用的函数；
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
	}
}

func main() {
	once()
}
