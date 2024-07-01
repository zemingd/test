package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var (
	// ReadString returns a WORD string.
	ReadString func() string
)

func init() {
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
}

func newReadString(ior io.Reader, sf bufio.SplitFunc) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+9)) // for Codeforces
	r.Split(sf)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

func readInt() int {
	n, _ := strconv.Atoi(ReadString())
	return n
}

// 10 11 12 => [10, 11, 12]
func readIntSlice(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = readInt()
	}
	return a
}

func get2byte(size int) [][]byte {
	a := make([][]byte, size)
	for i := 0; i < size; i++ {
		var low string
		fmt.Scan(&low)
		a[i] = []byte(low)
	}
	return a
}

func transpose(a [][]int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a[i][j], a[j][i] = a[j][i], a[i][j]
		}
	}
}

func fact(n, m int) int {
	res := 1
	for i := m + 1; i <= n; i++ {
		res *= i
	}
	return res
}

func perm(n, r int) int {
	if r > n {
		return 0
	}
	return fact(n, n-r)
}

func comb(n, m int) int {
	if m > n {
		return 0
	}
	return fact(n, n-m) / fact(m, 0)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func min(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m > integer {
			m = integer
		}
	}
	return m
}

func max(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m < integer {
			m = integer
		}
	}
	return m
}

func reverseInts(list []int) []int {
	n := len(list)
	tmp := make([]int, n)
	for i, l := range list {
		tmp[n-1-i] = l
	}
	return tmp
}

const mod = 1000000007

func main() {
	n, k := readInt(), readInt()
	as := readIntSlice(n)
	sort.Ints(as)
	res := solve(n, k, reverseInts(as))
	if res < 0 {
		res = mod + res
	}
	fmt.Println(res)
}

func multi(a, b int) int {
	res := (a * b) % mod
	return res
}
func solve(n, k int, as []int) int {
	if n == k {
		res := 1
		for _, a := range as {
			res = multi(res, a)
		}
		return res
	}
	maxAs := max(as...)
	if maxAs <= 0 {
		tmp := as
		if k%2 == 0 {
			tmp = reverseInts(as)
		}
		res := 1
		for i := 0; i < k; i++ {
			res = multi(res, tmp[i])
		}
		return res
	}

	var negativePart, zeroPart, positivePart []int
	for _, a := range as {
		if a > 0 {
			positivePart = append(positivePart, a)
		} else if a == 0 {
			zeroPart = append(zeroPart, a)
		} else {
			negativePart = append(negativePart, a)
		}
	}
	if len(positivePart)+len(negativePart) < k {
		return 0
	} else if len(positivePart)+len(negativePart) == k {
		if len(negativePart)%2 == 1 {
			return 0
		}
		res := 1
		for _, a := range positivePart {
			res = multi(res, a)
		}
		for _, a := range negativePart {
			res = multi(res, a)
		}
		return res
	}

	if len(negativePart) == 0 {
		res := 1
		for i := 0; i < k; i++ {
			res = multi(res, as[i])
		}
		return res
	}

	negativePart = reverseInts(negativePart)
	res := 1
	for i := 0; i < k-1; i++ {
		if len(negativePart) == 1 {
			res = multi(res, positivePart[0])
			positivePart = positivePart[1:]
		} else if len(positivePart) == 1 {
			res = multi(res, negativePart[0])
			negativePart = negativePart[1:]
		} else if positivePart[0] >= -negativePart[0] {
			res = multi(res, positivePart[0])
			positivePart = positivePart[1:]
		} else {
			res = multi(res, negativePart[0])
			negativePart = negativePart[1:]
		}
	}
	if res < 0 {
		res = multi(res, negativePart[0])
	} else {
		res = multi(res, positivePart[0])
	}
	return res
}
