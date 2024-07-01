package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	A := readInt()
	B := readInt()

	N := gcd(A, B)
	if N == 1 {
		fmt.Println(1)
		return
	}
	if N == 2 {
		fmt.Println(2)
		return
	}

	result := 1
	for x := 2; x < int(math.Sqrt(float64(N)))+1; x++ {
		if N%x == 0 {
			result++
			for N%x == 0 {
				N /= x
			}
		}
	}
	if result == 1 {
		result++
	}
	fmt.Println(result)
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
