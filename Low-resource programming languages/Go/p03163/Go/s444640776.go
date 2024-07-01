package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func chmax64(x *int64, y int64) bool {
	if *x < y {
		*x = y
		return true
	}
	return false
}

type DP struct {
	w int
	v int64
}

func main() {
	N := readi()
	W := readi()
	w := make([]int, N)
	v := make([]int64, N)
	for i := 0; i < N; i++ {
		w[i] = readi()
		v[i] = readi64()
	}

	dp := make([][]DP, N)
	dp[0] = make([]DP, N)
	for i := 0; i < N; i++ {
		if W >= w[i] {
			dp[0][i].w = w[i]
			dp[0][i].v = v[i]
		}
	}
	for i := 1; i < N; i++ {
		dp[i] = make([]DP, N)
		for ii := 0; ii < N; ii++ {
			for iii := 0; iii < N; iii++ {
				if iii != ii {
					if W >= w[ii]+dp[i-1][iii].w {
						chmax(&dp[i][ii].w, w[ii]+dp[i-1][iii].w)
						chmax64(&dp[i][ii].v, v[ii]+dp[i-1][iii].v)
					} else {
						chmax(&dp[i][ii].w, dp[i-1][iii].w)
						chmax64(&dp[i][ii].v, dp[i-1][iii].v)
					}
				}
			}
		}
	}

	var ans int64
	for i := range dp[N-1] {
		ans = max64(ans, dp[N-1][i].v)
	}

	fmt.Println(ans)
}

/*-------------------utilities-------------------*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func chmin(x *int, y int) bool {
	if *x > y {
		*x = y
		return true
	}
	return false
}

func chmax(x *int, y int) bool {
	if *x < y {
		*x = y
		return true
	}
	return false
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

/*-------------------init-------------------*/

const (
	INF = 1000000000000000000
	MOD = 1e9 + 7
)

var (
	readString func() string
)

func init() {
	readString = newReadString(os.Stdin)
}

/*-------------------inputs-------------------*/

func readi() int {
	return int(readi64())
}

func readi64() int64 {
	i, err := strconv.ParseInt(readString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readf() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}
