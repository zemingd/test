package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func solve(s string) int {
	if s == "SSS" {
		return 0
	}
	if s == "RSS" {
		return 1
	}
	if s == "SRS" {
		return 1
	}
	if s == "RRS" {
		return 2
	}
	if s == "SSR" {
		return 1
	}
	if s == "RSR" {
		return 1
	}
	if s == "SRR" {
		return 2
	}
	if s == "RRR" {
		return 3
	}
	return 0
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s := getString()
	out(solve(s))
	// for i := 0; i < 8; i++ {
	// 	s := ""
	// 	for j := 0; j < 3; j++ {
	// 		if i&(1<<j) != 0 {
	// 			s = s + "R"
	// 		} else {
	// 			s = s + "S"
	// 		}
	// 	}
	// 	out(s, solve(s))
	// }
}
