package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inf = 1 << 60

func solve() {
	d, g := getInt(), getInt()
	ps, cs := make([]int, d), make([]int, d)
	for i := 0; i < d; i++ {
		ps[i] = getInt()
		cs[i] = getInt()
	}

	ans := inf
	for bit := 0; bit < (1 << uint(d)); bit++ {
		score := 0
		cnt := 0
		restMax := -1
		for i := 0; i < d; i++ {
			if mask := (1 << uint(i)); (bit & mask) > 0 {
				score += (100*(i+1))*ps[i] + cs[i]
				cnt += ps[i]
			} else {
				restMax = i
			}
		}

		if diff := g - score; diff > 0 {
			if restMax == -1 {
				continue
			}

			s := 100 * (restMax + 1)
			add := (diff + s - 1) / s
			if add > ps[restMax] {
				continue
			}
			cnt += add
		}

		ans = min(ans, cnt)
	}

	fmt.Fprintln(wr, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// -----------------------------------------

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func main() {
	// buf := 200001
	// sc.Buffer(make([]byte, buf), buf)
	sc.Split(bufio.ScanWords)
	solve()
	wr.Flush()
}

func getInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func getIntSlice(n int) []int {
	is := make([]int, n)
	for i := range is {
		is[i] = getInt()
	}
	return is
}

func getString() string {
	sc.Scan()
	s := sc.Text()
	return s
}

func getRunes() []rune {
	return []rune(getString())
}
