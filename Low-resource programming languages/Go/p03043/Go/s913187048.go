package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	ret := 0
	for i := 1; i <= n; i++ {
		tmp := 1.0 / n
		now := i
		for now < k {
			now *= 2
			tmp /= 2
		}
		ret += tmp
	}
	fmt.Print(ret)

}
