package main

import "fmt"

func main() {
	var k, s int
	fmt.Scan(&k, &s)

	var cnt int
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z := s - x - y
			if 0 <= z && z <= k {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
