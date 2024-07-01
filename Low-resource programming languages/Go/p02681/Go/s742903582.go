package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	list := getStdinIntArr()
	A := list[0]
	B := list[1]
	C := list[2]
	K := list[3]

	if A - K > 0 {
		fmt.Printf("%d\n", A)
		return
	}
	K = -(A-K)

	if B - K > 0 {
		fmt.Printf("%d\n", A)
		return
	}
	K = -(B-K)

	if C - K > 0 {
		fmt.Printf("%d\n", A - C)
		return
	}

	fmt.Printf("%d\n", A + (C-K))
}

func getStdin() string {
	sc.Scan()
	return sc.Text()
}
func getStdinInt() int {
	str := getStdin()
	rtn, _ := strconv.Atoi(str)
	return rtn
}
func getStdinIntArr() []int {
	str := getStdin()
	list := strings.Split(str, " ")
	rtn := make([]int, len(list))
	for idx, val := range list {
		rtn[idx], _ = strconv.Atoi(val)
	}
	return rtn
}
