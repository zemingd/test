package main
 
import (
	"sort"
	"bufio"
	"fmt"
	"os"
	"strconv"
)
 
var sc, wr = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
 
func scanString() string { sc.Scan(); return sc.Text() }
func scanRunes() []rune { return []rune(scanString()) }
func scanInt() int { a,_ := strconv.Atoi(scanString()); return a }
func scanInt64() int64 { a,_ := strconv.ParseInt(scanString(),10,64); return a }
func scanFloat64() float64 { a,_ := strconv.ParseFloat(scanString(),64); return a }
 
func scanInts(n int) []int {
	res := make([]int, n); for i := 0; i < n; i++ { res[i] = scanInt() }; return res
}
 
func debug(a ...interface{}) { fmt.Fprintln(os.Stderr, a...) }
 
func abs(a int) int { if a<0 { return -a }; return a }
func min(a,b int) int { if a<b { return a }; return b }
func max(a,b int) int { if a>b { return a }; return b }
 
//•*¨*•.¸¸♪Main•*¨*•.¸¸♪(　-ω-)ノ　(　・ω・)
 
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)
 
	n := scanInt()
	h := scanInt()
 
	a := make([]int, n)
	maxa := 0
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		maxa = max(maxa, a[i])
		b[i] = scanInt()
	}

	sort.Ints(b)

	c := 0
	for i := 0; i < n; i++ {
		if maxa > b[n-1-i] {
			break
		}

		h -= b[n-1-i]
		c++

		if h <= 0 {
			h = 0
			break
		}
	}

	fmt.Fprintln(wr, c+(h+maxa-1)/maxa)
 
}
