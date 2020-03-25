package main

import "fmt"

func if_use() {
	num := 10
	if num%2 == 0 {
		fmt.Println("the number is even.")
	} else {
		fmt.Println("the number is odd.")
	}

	if num = 11; num > 10 {
		fmt.Println("the number bigger than ten.")
	}
}

func main() {
	if_use()
}
