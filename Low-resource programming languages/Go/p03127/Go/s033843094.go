package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	var min int
	fmt.Scan(&min)
	
	for i := 0; i < N-1; i++ {
		var a int
		fmt.Scan(&a)
		if a != min {
			if a % 2 == 1 || min % 2 == 1 {
				min = 1
				break
			} else {
                for ;; {
                  if min < a {
                    if a % min > 0 {
                        a = a % min
                    } else {
                        break
                    }
                  } else {
                    if min % a > 0 {
                        min = min % a
                    } else {
                        min = a
                        break
                    }
                  }
                }
			}
		}
	}
	fmt.Println(min)
}
