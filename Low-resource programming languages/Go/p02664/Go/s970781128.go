package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	t := sc.Text()
	a := strings.Split(t, "")
	if a[len(a)-1] == "?" { a[len(a)-1] = "D"}

	for i := len(a)-2; i >= 0; i-- {
		if a[i] != "?" { continue }
		if a[i-1] == "D" {
			a[i] = "P"
		} else {
			a[i] = "D"
		}
	}

	fmt.Println(strings.Join(a, ""))
}

func slicePrint(s []int) {
	for i, v := range s {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%v", v)
	}
	fmt.Println("")
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
