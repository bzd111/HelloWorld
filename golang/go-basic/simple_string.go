package main

import (
	"fmt"
)

func main() {

	s := "HelloWorld"
	// runtime.stringtoslicebyte
	bytes := []byte(s)
	bytes[0] = 'h'
	fmt.Printf("byte: %v\n", bytes)

	// runtime.slicebytetostring
	s = string(bytes)
	fmt.Println("string:", s)
}
