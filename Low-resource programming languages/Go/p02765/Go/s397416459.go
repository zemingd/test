package main

import "fmt"

func main() {
	N, R := read(), read()

	if N >= 10 {
		println(R)
		return
	} else {
		println(R + 1000 - 100*N)
	}
}

func read() int {
	var v int
	fmt.Scan(&v)
	return v
}
