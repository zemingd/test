package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	if IsPrime(n) {
		fmt.Println(1)
		return
	}

	count := 0
	for _, cnt := range PrimeFactor(n) {
		for i := 1; i <= cnt; i++ {
			count++
			cnt -= i
		}
	}

	fmt.Println(count)
}

// generated by https://github.com/murosan/gollect

func PrimeFactor(n int) map[int]int {
	m := make(map[int]int)
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	return m
}

func IsPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n < 2 || n%2 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := int(3); i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
