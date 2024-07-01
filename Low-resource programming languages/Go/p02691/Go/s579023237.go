package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	l := make([]int, N)
	r := make([]int, N)
	for i := 0; i < N; i++ {
		if i-a[i] > 0 {
			l[i-a[i]]++
		}
		if i+a[i] < N {
			r[i+a[i]]++
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
		ans += r[i] * l[i]
	}

	// out(r)
	// out(l)
	out(ans)
	// ans := 0
	// for i := 0; i < N; i++ {
	// 	for j := i + a[i]; j < N; j++ {
	// 		if a[i]+a[j] == j-i {
	// 			ans++
	// 		}
	// 	}
	// }
	// out(ans)
}
