package main

import "fmt"

func main(){
	var n,a,b int
	fmt.Scan(&n,&a,&b)
	if n*a>b{
		fmt.Printf("%d\n",b)
	} else{
		fmt.Printf("%d\n",n*a)
	}

}