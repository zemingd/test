package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	fmt.Println(answer162B(n))
}

func answer162B(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		if !(i%3 == 0 || i%5 == 0) {
			sum += i
		}
	}

	return sum
}
