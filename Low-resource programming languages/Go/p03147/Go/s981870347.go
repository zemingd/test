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

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	h := make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = getInt()
	}

	ans := 0
	for {
		l := 0
		for i := 0; i < N; i++ {
			if h[i] != 0 {
				break
			}
			l++
		}
		mi := 10000
		r := l
		for i := l; i < N; i++ {
			if h[i] == 0 {
				break
			}
			mi = min(mi, h[i])
			r++
		}
		// out(l, r, h)
		if l == r {
			break
		}
		for i := l; i < r; i++ {
			h[i] -= mi
		}
		ans += mi
	}
	out(ans)
}
