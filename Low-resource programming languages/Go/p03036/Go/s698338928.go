package main

import "fmt"

func main() {
	var r, d, x int
	fmt.Scan(&r, &d, &x)
	for i := 0; i < 10; i++ {
		x = x*r - d
		fmt.Println(x)
	}
}
