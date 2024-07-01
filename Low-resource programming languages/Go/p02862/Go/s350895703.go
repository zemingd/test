package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	M = 1000000007
)

var (
	fac  []int
	ifac []int
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func mpow(x int, n int) int {
	ans := 1
	for n != 0 {
		if n&1 == 1 {
			ans = ans * x % M
		}
		x = x * x % M
		n = n >> 1
	}
	return ans
}

func comb(a int, b int) int {
	if a == 0 && b == 0 {
		return 1
	}
	if (a-b < 0) || (a < 0) {
		return 0
	}
	fmt.Println(a - b)
	fmt.Println(b)
	tmp := ifac[a-b] * ifac[b] % M
	return tmp * fac[a] % M
}

func main() {
	X := readInt()
	Y := readInt()

	if (2*Y-X)%3 != 0 || (2*X-Y)%3 != 0 {
		fmt.Println(0)
		return
	}
	if 2*Y-X < 0 || 2*X-Y < 0 {
		fmt.Println(0)
		return
	}

	fac = make([]int, 3000001)
	ifac = make([]int, 3000001)

	a := (2*Y - X) / 3
	b := (2*X - Y) / 3

	fac[0] = 1
	ifac[0] = 1
	for i := 0; i < 2000000; i++ {
		fac[i+1] = fac[i] * (i + 1) % M
		ifac[i+1] = ifac[i] * mpow(i+1, M-2) % M
	}

	fmt.Println(comb(a+b, min(a, b)) % M)
}

const (
	ioBufferSize = 1 * 1024 * 1024 // 1 MB
)

var stdinScanner = func() *bufio.Scanner {
	result := bufio.NewScanner(os.Stdin)
	result.Buffer(make([]byte, ioBufferSize), ioBufferSize)
	result.Split(bufio.ScanWords)
	return result
}()

func readString() string {
	stdinScanner.Scan()
	return stdinScanner.Text()
}

func readInt() int {
	result, err := strconv.Atoi(readString())
	if err != nil {
		panic(err)
	}
	return result
}
