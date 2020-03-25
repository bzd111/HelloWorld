package main

import (
	"fmt"
	"time"
)

func notInGoroutine() {
	// 只会触发当前Goroutine的延迟函数
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)
}

func recoverIndefer() {
	defer func() {
		// recover需要在defer中调用
		if err := recover(); err != nil {
			println("recover success")
		}
	}()
	panic("panicing...")
}

func nestPanic() {
	// 嵌套panic，最外层的defer也会执行
	// defer println("in main")
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()
	panic("panic")
}

func main() {
	// notInGoroutine()
	// recoverIndefer()
	nestPanic()
}
