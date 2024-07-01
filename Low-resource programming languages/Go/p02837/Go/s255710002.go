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

type rule struct {
	x, y int
}

// N :
var N int
var a []int
var r [][]rule

func check(n int) int {
	b := make([]int, N)
	idx := 0
	for n > 0 {
		b[idx] = n % 2
		n /= 2
		idx++
	}

	ret := 0
	for i := 0; i < N; i++ {
		// out(i, N)
		if b[i] == 1 {
			ret++
			// out(i, r[i])
			for _, e := range r[i] {
				if b[e.x] != e.y {
					return -1
				}
			}
		}
	}
	// out(b)
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	a = make([]int, N)
	r = make([][]rule, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
		for j := 0; j < a[i]; j++ {
			x, y := getInt()-1, getInt()
			r[i] = append(r[i], rule{x, y})
		}
	}

	ans := 0
	n := 1 << uint(N)
	for i := 0; i < n; i++ {
		res := check(i)
		ans = max(ans, res)
	}
	out(ans)
}
