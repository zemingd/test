package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var buff []byte

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextFloat64() float64 {
	sc.Scan()
	f, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func nextInts(n int) (r []int) {
	r = make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = nextInt()
	}
	return r
}

var dy = []int{0, 1, 0, -1}
var dx = []int{1, 0, -1, 0}
var MAX = math.MaxInt64

func maxInt(a ...int) int {
	r := a[0]
	for i := 0; i < len(a); i++ {
		if r < a[i] {
			r = a[i]
		}
	}
	return r
}
func minInt(a ...int) int {
	r := a[0]
	for i := 0; i < len(a); i++ {
		if r > a[i] {
			r = a[i]
		}
	}
	return r
}
func sum(a []int) (r int) {
	for i := range a {
		r += a[i]
	}
	return r
}
func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Pair struct {
	a, b int
}

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}
func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pairs) Less(i, j int) bool {
	if p[i].a == p[j].a {
		return p[i].b < p[j].b
	}
	return p[i].a < p[j].a
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buff, bufio.MaxScanTokenSize*1024)
	log.SetFlags(log.Lshortfile)
}

var inf = math.MaxInt64

type Point struct {
	x, y int
}

var H, W, K int
var i, j int

func main() {
	// s := time.Now()
	H = nextInt()
	W = nextInt()
	K = nextInt()
	var start, goal Point
	start.y = nextInt() - 1
	start.x = nextInt() - 1
	goal.y = nextInt() - 1
	goal.x = nextInt() - 1
	grid := make([]string, H)
	for i = 0; i < H; i++ {
		grid[i] = nextString()
	}
	step := make([][]int, H)
	for i = 0; i < H; i++ {
		step[i] = make([]int, W)
	}
	for i = 0; i < H; i++ {
		for j = 0; j < W; j++ {
			step[i][j] = inf
		}
	}
	step[start.y][start.x] = 0
	q := make([]*Point, 0)
	q = append(q, &start)
	idx := 0
	var next Point
	for idx < len(q) {
		now := q[idx]
		idx++
		for d := 0; d < 4; d++ {
			var k int
			for k = 1; k <= K; k++ {
				next.x = now.x + dx[d]*k
				next.y = now.y + dy[d]*k
				if next.x < 0 || next.x >= W || next.y < 0 || next.y >= H {
					break
				}
				if grid[next.y][next.x] == '@' {
					break
				}
				//if step[next.y][next.x] <= step[now.y][now.x] {
				//	break
				//}
				if step[next.y][next.x] <= step[now.y][now.x]+1 {
					continue
				}
				step[next.y][next.x] = step[now.y][now.x] + 1
				q = append(q, &Point{x: next.x, y: next.y})
			}
		}
		// elapsed := time.Since(s)
		// if elapsed > 2900*time.Millisecond {
		// 	break
		// }
	}
	if step[goal.y][goal.x] == inf {
		fmt.Println("-1")
	} else {
		fmt.Println(step[goal.y][goal.x])
	}
}
