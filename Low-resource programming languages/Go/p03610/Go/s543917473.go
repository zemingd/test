package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < len(s); i += 2 {
		fmt.Printf("%s", string(s[i]))
	}
	fmt.Println()
}
