package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var (
	readString func() string
	readBytes  func() []byte
	stdout     *bufio.Writer
)

func init() {
	readString, readBytes = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) (func() string, func() []byte) {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), math.MaxInt32)
	r.Split(bufio.ScanWords)

	f1 := func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
	f2 := func() []byte {
		return []byte(f1())
	}
	return f1, f2
}

func readInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func printf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(stdout, f, args...)
}

func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(stdout, args...)
}

func eprintln(args ...interface{}) (int, error) {
	return fmt.Fprintln(os.Stderr, args...)
}

func eprintf(f string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stderr, f, args...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func max64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcd64(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd64(b, a%b)
}

// readString() string
// readInt() int
// readInt64() int64
// readFloat64() float64

// -----------------------------------------------------------------------------

func rev(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	defer stdout.Flush()
	n := readInt()
	a := make([]int, n)
	for i := range a {
		a[i] = readInt()
	}
	c := make([][]int, 2)
	c[0] = make([]int, 0, len(a)/2+1)
	c[1] = make([]int, 0, len(a)/2+1)
	for i, v := range a {
		c[i%2] = append(c[i%2], v)
	}
	b := make([]int, 0, len(a))
	rev(c[0])
	for i := range c {
		for _, v := range c[i] {
			b = append(b, v)
		}
	}
	if len(a)%2 == 0 {
		rev(b)
	}

	sp := ""
	for _, v := range b {
		printf("%s%d", sp, v)
		sp = " "
	}
	println()
}

// -----------------------------------------------------------------------------
