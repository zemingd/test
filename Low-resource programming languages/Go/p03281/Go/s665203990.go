package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var ans int
	for i := 1; i <= n; i+=2 {
		count := 0
		for j := 1; j <= n; j++ {
			if i % j == 0 {
				count++
			}
			if count == 8 {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
