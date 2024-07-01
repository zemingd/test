package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func oneInt() int {
	var a int
	fmt.Scan(&a)
	return a
}

func oneInt64() int64 {
	var a int64
	fmt.Scan(&a)
	return a
}

func oneStr() string {
	var a string
	fmt.Scan(&a)
	return a
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func scanNums() (nums []int) {
	s := nextLine()
	numStr := strings.Split(s, " ")

	for _, n := range numStr {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
	}
	return nums
}

func scanStrings() (strs []string) {
	s := nextLine()
	list := strings.Split(s, " ")
	strs = append(strs, list...)
	return strs
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}

func strSearch(a []string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
}

func sortAsc(a []int) []int {
	sort.Ints(a)
	return a
}

func sortDesc(a []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	return a * b / gcd(a, b)
}

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

func homogeneous(n int, k int) int {
	return combination(n+k-1, k)
}

type key struct {
	w uint
	v uint
}

type dpKey struct {
	i int
	w uint
}

var table map[int]key
var dp [][]uint

func main() {
	N := oneInt()
	W := uint(oneInt())

	table = make(map[int]key)
	dp = make([][]uint, N+1)
	for index := 0; index <= N; index++ {
		dp[index] = make([]uint, 110000)
	}

	for index := 0; index < N; index++ {
		table[index] = key{uint(oneInt()), uint(oneInt())}
	}

	for index := 0; index < N; index++ {
		var windex uint
		for windex = 0; windex <= W; windex++ {
			if table[index].w <= windex { // 対象の荷物が入るかチェック
				temp := table[index].v + dp[index][windex- table[index].w] 
				chmax(&dp[index+1][windex], temp)
			} //入るなら入れてみて、余った空間に入る最大の勝ちのものをDPテーブルから持ってくる
			chmax(&dp[index+1][windex], dp[index][windex])
		}
	}
	fmt.Print(dp[N][W])
}

func chmax(a *uint, b uint) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
