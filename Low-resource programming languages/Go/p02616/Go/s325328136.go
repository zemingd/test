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

const mod = int(1e9 + 7)

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	a := getInts(N)

	if N == K {
		ans := 1
		for i := 0; i < K; i++ {
			ans *= a[i]
			ans %= mod
			if ans < 0 {
				ans += mod
			}
		}
		out(ans)
		return
	}

	plus := make([]int, 0)
	minus := make([]int, 0)
	for i := 0; i < N; i++ {
		if a[i] >= 0 {
			plus = append(plus, a[i])
		} else {
			minus = append(minus, -a[i])
		}
	}
	if len(minus) == N && K%2 == 1 {
		ans := 1
		for i := 0; i < K; i++ {
			ans *= a[i]
			ans %= mod
			if ans < 0 {
				ans += mod
			}
		}
		out(ans)
		return
	}

	sort.Sort(sort.Reverse(sort.IntSlice(plus)))
	sort.Sort(sort.Reverse(sort.IntSlice(minus)))

	posP := 0
	posM := 0
	for i := 0; i < K; i++ {
		if plus[posP] >= minus[posM] {
			posP++
		} else {
			posM++
		}
	}

	if posM%2 == 1 {
		if len(minus) == posM {
			posM--
			posP++
		} else if len(plus) == posP {
			posP--
			posM++
		} else {
			if minus[posM-1]*minus[posM] > plus[posP-1]*plus[posP] {
				posM++
				posP--
			} else {
				posM--
				posP++
			}
		}
	}

	ans := 1
	for i := 0; i < posP; i++ {
		ans *= plus[i]
		ans %= mod
	}
	for i := 0; i < posM; i++ {
		ans *= minus[i]
		ans %= mod
	}
	out(ans)
}
