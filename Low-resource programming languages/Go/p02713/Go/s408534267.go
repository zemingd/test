package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func next() string {
	sc.Scan()
	return sc.Text()
}
func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}
func nextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = nextInt()
	}
	return ret
}
func nextInt64() int64 {
	a, _ := strconv.ParseInt(next(), 10, 64)
	return a
}
func nextFloat64() float64 {
	a, _ := strconv.ParseFloat(next(), 64)
	return a
}

//var done []bool = make([]bool, 10)
//var memo []int = make([]int, 10)

type Group struct {
	Depth int
	Value string
}

type Job struct {
	No     int
	S      int
	E      int
	Worker string
}

type Res struct {
	No int
}

func Abs(a int64, b int64) int64 {
	res := a - b
	if a-b < 0 {
		res = -res
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	sc.Split(bufio.ScanWords)
	defer out.Flush() // !!!!coution!!!! you must use Fprint(out, ) not Print()
	n := nextInt()
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				res += gcd(gcd(i, j), k)
			}
		}
	}
	fmt.Fprintln(out, fmt.Sprint(res))
}
