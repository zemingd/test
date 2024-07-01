package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type segmentTree struct {
	data   []int
	offset int
}

func newSegmentTree(n int) segmentTree {
	var result segmentTree
	t := 1
	for t < n {
		t *= 2
	}
	result.offset = t - 1
	result.data = make([]int, 2*t-1)
	return result
}

func (st segmentTree) update(index, value int) {
	i := st.offset + index
	st.data[i] = value
	for i >= 1 {
		i = (i - 1) / 2
		st.data[i] = gcd(st.data[i*2+1], st.data[i*2+2])
	}
}

func (st segmentTree) _query(start, stop, index, left, right int) int {
	if right <= start || stop <= left {
		return 0
	}
	if start <= left && right <= stop {
		return st.data[index]
	}
	a := st._query(start, stop, index*2+1, left, (left+right)/2)
	b := st._query(start, stop, index*2+2, (left+right)/2, right)
	return gcd(a, b)
}

func (st segmentTree) query(start, stop int) int {
	return st._query(start, stop, 0, 0, st.offset+1)
}

func main() {
	N := readInt()
	A := readInts(N)

	st := newSegmentTree(N)
	for i, v := range A {
		st.update(i, v)
	}

	result := st.query(1, N)
	for i := 1; i < N-1; i++ {
		result = max(result, gcd(st.query(0, i), st.query(i+1, N)))
	}
	result = max(result, st.query(0, N-1))
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

func readInts(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = readInt()
	}
	return result
}

func printIntln(v ...int) {
	b := make([]byte, 0, 4096)
	for i := 0; i < len(v)-1; i++ {
		b = append(b, strconv.Itoa(v[i])...)
		b = append(b, " "...)
	}
	b = append(b, strconv.Itoa(v[len(v)-1])...)
	fmt.Println(string(b))
}
