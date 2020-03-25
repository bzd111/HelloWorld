package main

import "fmt"

// select 能在 Channel 上进行非阻塞的收发操作；
// select 在遇到多个 Channel 同时响应时会随机挑选 case 执行；

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			x := <-c
			fmt.Println("x: ", x)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
