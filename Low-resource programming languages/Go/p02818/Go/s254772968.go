package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)
	for i:=0; i<k; i++{
		if a>=1 {
			a--
		}else if b >=1{
			b--
		}

		if a==0 && b == 0{
			break
		}
	}
	fmt.Println(a, b)
}