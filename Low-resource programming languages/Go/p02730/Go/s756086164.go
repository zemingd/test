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

func main() {

	S := oneStr()
	var last int
	var start int
	if (len(S)-1)%2 == 0 {
		last = (len(S) - 1) / 2
	} else {
		last = (len(S) - 1) / 2
	}

	if (len(S)+3)%2 == 0 {
		start = (len(S)+3)/2 - 1
	} else {
		start = (len(S) + 3) / 2
	}

	if isKaibun(S[:last]) {
		if isKaibun(S[start:]) {
			fmt.Print("Yes")
			return
		}
	}
	fmt.Print("No")
}

func isKaibun(s string) bool {
	if len(s) == 1 {
		return true
	}
	fmt.Print("%s\n", s)
	for i := 0; i != len(s)-i-1; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func exp(val int) float64 {
	var temp float64
	for index := 1; index <= val; index++ {
		temp += float64(index)
	}
	return temp / float64(val)
}

func searchAbove(slice []int, target int, index int) int {
	for i := index + 1; i < len(slice); i++ {
		v := slice[i]
		if v >= target {
			return i
		}
	}
	return -1
}

var sc = bufio.NewScanner(os.Stdin)

func oneInt() int {
	var a int
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
