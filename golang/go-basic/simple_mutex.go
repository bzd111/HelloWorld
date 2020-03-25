package main

import (
	"fmt"
	"sync"
)

func notCopy() {
	wg := sync.Mutex{}
	yawg := wg
	fmt.Println(wg, yawg)
}
func main() {
	notCopy()
	// wp := sync.WaitGroup{}
	// yawg := wp
	// fmt.Println(wp, yawg)
}
