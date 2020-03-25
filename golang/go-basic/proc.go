package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	yawg := wg
	fmt.Println(wg, yawg)
}
