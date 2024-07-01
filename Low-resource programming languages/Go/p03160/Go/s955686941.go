package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	initialBufSize = 1024
	maxBufSize     = 1e6
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriterSize(os.Stdout, 1000000)
)

func main() {
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	n := scanInt()
	h := scanInts(n)
	dp := make([]int, n)
	min := func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i < n; i++ {
		dp[i] = min(dp[i], dp[i-1]+abs(h[i]-h[i-1]))
		if i > 1 {
			dp[i] = min(dp[i], dp[i-2]+abs(h[i]-h[i-2]))
		}
	}

	fmt.Println(dp[n-1])
}

func scanInt() int {
	i, _ := strconv.Atoi(scanString())
	return i
}

func scanInts(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = scanInt()
	}
	return ints
}

func scanFloat64() float64 {
	f, _ := strconv.ParseFloat(scanString(), 64)
	return f
}

func scanFloat64s(size int) []float64 {
	f := make([]float64, size)
	for i := 0; i < size; i++ {
		f[i] = scanFloat64()
	}
	return f
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanStrings(size int) []string {
	s := make([]string, size)
	for i := 0; i < size; i++ {
		s[i] = scanString()
	}
	return s
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// greatest common divisor(最大公約数)
func gcd(a, b int) int {
	for {
		mod := a % b
		if mod == 0 {
			return b
		}
		a, b = b, mod
	}
}

// latest common multiple(最小公倍数)
func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}

func fact(n int) int {
	ret := 1
	for i := 1; i <= n; i++ {
		ret *= i
	}
	return ret
}

func perm(n, r int) int {
	ret := 1
	for i := 0; i < r; i++ {
		ret *= n
		n--
	}
	return ret
}

func isPrimeNum(n int) bool {

	if n <= 1 {
		return false
	}

	l := int(math.Sqrt(float64(n)) + 1.)
	for i := 2; i < l; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
