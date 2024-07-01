package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)
	count := 0
	tmp := 0
	for i := 0; i < N-1; i++ {
		l := string(S[i])
		r := string(S[i+1])
		if l == "#" && r == "#" {
			tmp++
		}
		if l == "#" && r == "." {
			count = tmp + 1
			tmp = 0
		}
	}

	fmt.Println(count)
}