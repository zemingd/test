package main

import (
	"fmt"
	//"math"
	//"math"
	//"strings"
)

func main() {
	var a,b int
	fmt.Scan(&a,&b)
	if b % a == 0 {
		fmt.Println(a+b)
	}else{
		fmt.Println(b-a)
	}
}