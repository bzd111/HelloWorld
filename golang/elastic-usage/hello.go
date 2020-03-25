package main

import (
	"fmt"
)

func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		k := i
		out = append(out, &k)
	}
	// fmt.Println("Values:", *out[0], *out[1], *out[2])
	// fmt.Println("Addresses:", out[0], out[1], out[2])
	var out1 [][]int
	for _, i := range [][1]int{{1}, {2}, {3}} {
		c := i
		out1 = append(out1, c[:])
	}
	fmt.Println("Values:", out1)
}
