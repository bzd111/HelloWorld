package main

import (
	"fmt"
)

type ByteSize float64

const (
	_           = iota // 从零开始，在这里需要忽略
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("MB: %.2f\n", MB)
	fmt.Printf("YB: %.2f\n", YB)
}
