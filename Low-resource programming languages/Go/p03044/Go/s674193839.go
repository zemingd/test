package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	readString func() string
	readBytes  func() []byte
)

func init() {
	readString, readBytes = newReadString(os.Stdin)
}

func newReadString(ior io.Reader) (func() string, func() []byte) {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)

	f1 := func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
	f2 := func() []byte {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Bytes()
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func absInt64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func sum(a []int) int {
	var ret int
	for i := range a {
		ret += a[i]
	}
	return ret
}

func sumInt64(a []int64) int64 {
	var ret int64
	for i := range a {
		ret += a[i]
	}
	return ret
}

func sumFloat64(a []float64) float64 {
	var ret float64
	for i := range a {
		ret += a[i]
	}
	return ret
}

// sort ------------------------------------------------------------

type int64Array []int64

func (s int64Array) Len() int           { return len(s) }
func (s int64Array) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s int64Array) Less(i, j int) bool { return s[i] < s[j] }

type xxx struct {
	x int
}

type sortArray []xxx

func (s sortArray) Len() int           { return len(s) }
func (s sortArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortArray) Less(i, j int) bool { return s[i].x < s[j].x }

// -----------------------------------------------------------------

func dfs(idx int, d int64) {
	if d%2 == 0 {
		color[idx-1] = 0
	} else {
		color[idx-1] = 1
	}
	for next, w := range tree[idx] {
		if color[next-1] != -1 {
			continue
		}
		dfs(next, d+w)
	}
}

var tree map[int]map[int]int64
var color []int

func main() {
	n := readInt()
	tree = make(map[int]map[int]int64)
	for i := 0; i < n-1; i++ {
		u := readInt()
		v := readInt()
		w := readInt64()
		if _, ok := tree[u]; !ok {
			tree[u] = make(map[int]int64)
		}
		if _, ok := tree[v]; !ok {
			tree[v] = make(map[int]int64)
		}
		tree[u][v] = w
		tree[v][u] = w
	}
	color = make([]int, n)
	for i := range color {
		color[i] = -1
	}

	dfs(1, 0)

	for i := range color {
		fmt.Println(color[i])
	}
}
