package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	if a < b && b < c {
		fmt.Print("Yes\n")
	} else {
		fmt.Print("No\n")
	}

}
