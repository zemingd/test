package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	for i := 0; i < len(s); i++ {
		if s == t {
			fmt.Println("Yes")
			return
		}
		s = s[1:] + s[:1]
	}
	fmt.Println("No")
}
