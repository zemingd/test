package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*-----------Main-----------*/

func main() {
	n, k := readI(), readI()
	town := make([]int, n+1)
	town[0] = 1
	for i := 1; i <= n; i++ {
		town[i] = readI()
	}

	visited := make([]bool, n+1)
	seq := make([]int, 0, n+1)
	now := 0
	next := 1
	var loopStart, loopEnd int
	for step := 0; step <= n; step++ {
		next = town[now]
		seq = append(seq, now)
		visited[now] = true
		if visited[next] {
			loopStart = search(seq, next)
			loopEnd = step
			break
		}

		now = next
	}

	loopLen := loopEnd - loopStart + 1

	if k <= loopEnd {
		fmt.Println(seq[k-1])
	} else {
		fmt.Println(seq[loopStart-1+(k-(loopStart-1))%loopLen+1])
	}

}

func search(a []int, b int) int {
	l := len(a)
	for i := 0; i < l; i++ {
		if a[i] == b {
			return i
		}
	}
	return -1
}

/*-----------Utilities-----------*/

var sc = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	const maxBuf = 200100
	var buf []byte = make([]byte, maxBuf)
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBuf)
}

func debug() {
	testFile, err := os.Open("./test/sample-2.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error : There is no testfile.")
		os.Exit(1)
	}
	sc = bufio.NewScanner(testFile)
}

func readS() string {
	sc.Scan()
	return sc.Text()
}

func readSs(a int) []string {
	b := make([]string, a)
	for i := 0; i < a; i++ {
		b[i] = readS()
	}
	return b
}

func readI() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func readIs(a int) []int {
	b := make([]int, a)
	for i := 0; i < a; i++ {
		b[i] = readI()
	}
	return b
}

func readF() float64 {
	sc.Scan()
	r, _ := strconv.ParseFloat(sc.Text(), 64)
	return r
}

func readFs(a int) []float64 {
	b := make([]float64, a)
	for i := 0; i < a; i++ {
		b[i] = readF()
	}
	return b
}

func max(a []int) int {
	l := len(a)
	m := a[0]
	for i := 1; i < l; i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

func min(a []int) int {
	l := len(a)
	m := a[0]
	for i := 1; i < l; i++ {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
}

func sortI(a []int) []int {
	sort.Ints(a)
	return a
}

func reverse(a []int) []int {
	l := len(a)
	r := make([]int, l)
	for i := 0; i < l; i++ {
		r[i] = a[l-i-1]
	}
	return r
}

func countI(a []int, b int) int {
	l := len(a)
	var c int
	for i := 0; i < l; i++ {
		if a[i] == b {
			c++
		}
	}
	return c
}
