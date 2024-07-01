package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

type trees struct {
	Right []int
	Left  []int
}

func (tree *trees) newUTree(N int) *trees {
	u := &trees{Right: make([]int, 0), Left: make([]int, 0)}
	return u
}

func (tree *trees) addRight(e int) {
	tree.Right = append(tree.Right, e)
}

func (tree *trees) addLeft(e int) {
	tree.Right = append(tree.Right, e)
}

func (tree *trees) addBoth(e int) {
	tree.Right = append(tree.Right, e)
	tree.Left = append(tree.Left, e)
}

func main() {
	N, X, Y := sc.nextInt(), sc.nextInt(), sc.nextInt()
	var tree *trees
	tree = tree.newUTree(N)

	for i := 1; i < N+1; i++ {
		if (i) > X && (i) < Y {
			tree.addRight(i)
		} else {
			tree.addBoth(i)
		}
	}
	cMap := make(map[Pair]int)
	ansMap := make(map[int]int)
	for i := 0; i < len(tree.Right); i++ {
		for j := i + 1; j < len(tree.Right); j++ {
			cMap[Pair{tree.Right[j], tree.Right[i]}] = j - i
			cMap[Pair{tree.Right[i], tree.Right[j]}] = j - i
		}
	}
	for i := 0; i < len(tree.Left); i++ {
		for j := i + 1; j < len(tree.Left); j++ {
			cMap[Pair{tree.Left[j], tree.Left[i]}] = j - i
			cMap[Pair{tree.Left[i], tree.Left[j]}] = j - i
			// cMap[tree.Left[j]-tree.Left[i]]++
		}
	}
	// for i := 1; i < N; i++ {
	// 	fmt.Println(cMap[i])
	// }
	for _, v := range cMap {
		ansMap[v]++
	}
	for i := 1; i < N; i++ {
		fmt.Println(ansMap[i] / 2)
	}
}

var (
	sc scanner
)

func init() {
	sc = scanner{
		buf: make([]string, 0, 0),
		cur: 0,
		r:   bufio.NewReader(os.Stdin),
	}
}

type scanner struct {
	buf []string
	cur int
	r   *bufio.Reader
}

func (s *scanner) readln() {
	rbuf := make([]byte, 0, 0)
	for {
		line, prefix, err := s.r.ReadLine()
		if err != nil {
			panic(err)
		}
		rbuf = append(rbuf, line...)
		if prefix == false {
			break
		}
	}
	s.cur = 0
	s.buf = strings.Split(*(*string)(unsafe.Pointer(&rbuf)), " ")
}
func (s *scanner) isFull() bool {
	if s.cur == len(s.buf) {
		return true
	}
	return false
}
func (s *scanner) resetCur() {
	s.cur = 0
}
func (s *scanner) next() string {
	if s.cur == 0 {
		s.readln()
	}
	res := s.buf[s.cur]
	s.cur++
	if s.isFull() {
		s.resetCur()
	}
	return res
}
func (s *scanner) nexts() []string {
	s.readln()
	s.resetCur()
	return s.buf
}
func (s *scanner) nextInt() int {
	if s.cur == 0 {
		s.readln()
	}
	res, _ := strconv.Atoi(s.buf[s.cur])
	s.cur++
	if s.isFull() {
		s.resetCur()
	}
	return res
}
func (s *scanner) nextInts() []int {
	s.readln()
	res := make([]int, len(s.buf))
	for i := range s.buf {
		res[i], _ = strconv.Atoi(s.buf[i])
	}
	s.resetCur()
	return res
}
func (s *scanner) nextFloat() float64 {
	if s.cur == 0 {
		s.readln()
	}
	res, _ := strconv.ParseFloat(s.buf[s.cur],
		64)
	s.cur++
	if s.isFull() {
		s.resetCur()
	}
	return res
}
func (s *scanner) nextFloats() []float64 {
	s.readln()
	res := make([]float64, len(s.buf))
	for i := range s.buf {
		res[i], _ = strconv.ParseFloat(s.buf[i],
			64)
	}
	s.resetCur()
	return res
}
func digits(x int) int {
	return len(strconv.Itoa(x))
}
func powInt(x, p int) (result int) {
	result = 1
	for i := 0; i < p; i++ {
		result *= x
	}
	return
}
func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
func abs(x int) int {
	return int(math.Abs(float64(x)))
}
func varDump(value ...interface{}) {
	for _, v := range value {
		fmt.Fprintf(os.Stderr, "%#v\n", v)
	}
}
func yes() {
	fmt.Println("Yes")
}
func no() {
	fmt.Println("No")
}

// Pair a
type Pair struct {
	first  int
	second int
}

// Queue a
type Queue struct {
	v []Pair
}

// Push for queue
func (q *Queue) Push(v Pair) {
	q.v = append(q.v, v)
}

// Pop for queue
func (q *Queue) Pop() Pair {
	r := q.v[0]
	q.v = q.v[1:]
	return r
}

// Front for queue
func (q Queue) Front() Pair {
	return q.v[0]
}

// Empty return true if empty
func (q Queue) Empty() bool {
	return len(q.v) == 0
}
