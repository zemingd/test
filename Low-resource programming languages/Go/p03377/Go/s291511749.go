package main

import "fmt"

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)
	if x < a || (a + b) < x {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
