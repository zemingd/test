package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a*b%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
	return
}
