package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		if p[i] < p[i-1] && p[i] >= p[i+1] {
			ans++
			continue
		}
		if p[i] < p[i+1] && p[i] >= p[i-1] {
			ans++
			continue
		}
	}
	fmt.Println(ans)
}
