package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n - 1; i++ {
		sum = sum + i + 1
	}

	fmt.Println(sum)
}
