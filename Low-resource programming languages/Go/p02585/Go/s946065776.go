package main

import (
	"bufio"
	"fmt"
	"math"
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

const inf = math.MaxUint64 >> 1

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, K := getInt(), getInt()
	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getInt() - 1
	}
	c := getInts(N)
	ans := -inf
	for i := 0; i < N; i++ {
		used := make([]bool, N)
		score := make([]int, 0)
		ps := 0
		next := p[i]
		for used[next] != true {
			// out(next)
			used[next] = true
			score = append(score, ps+c[next])
			ps = ps + c[next]
			next = p[next]
		}
		n := K / len(score)
		m := K % len(score)
		var sum int
		if score[len(score)-1] >= 0 {
			sum = max(0, score[len(score)-1]*n)
			ma := 0
			for j := 0; j < m; j++ {
				ma = max(ma, score[j])
			}
			// out(sum, ma)
			sum += ma
		} else {
			ma := -inf
			l := min(len(score), K)
			for j := 0; j < l; j++ {
				ma = max(ma, score[j])
			}
			// out(score, ma)
			sum = ma
		}
		ans = max(ans, sum)
	}
	out(ans)
}
