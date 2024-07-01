package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := NewScanner()

	n := sc.ReadInt()
	m := sc.ReadInt()
	q := sc.ReadInt()
	abcds := make([]ABCD, q)
	for i := 0; i < q; i++ {
		a := sc.ReadInt()
		b := sc.ReadInt()
		c := sc.ReadInt()
		d := sc.ReadInt()
		abcds[i] = ABCD{a, b, c, d}
	}

	cur := make(SeqA, n)
	for i := 0; i < n-1; i++ {
		// the last is 0
		cur[i] = 1
	}

	maxScore := 0
	for {
		var ok bool
		ok, cur = nextSeqA(n, m, cur)
		if !ok {
			break
		}

		score := computeScore(abcds, cur)
		if score > maxScore {
			maxScore = score
		}
	}

	fmt.Println(maxScore)
}

func nextSeqA(n int, m int, cur SeqA) (bool, SeqA) {
	carryDigitCount := 0
	for i := n - 1; i >= 0; i-- {
		if cur[i] == m {
			carryDigitCount++
		} else {
			break
		}
	}

	if carryDigitCount == 0 {
		cur[n-1]++
		return true, cur
	} else if carryDigitCount == n {
		return false, nil
	} else {
		min := cur[n-1-carryDigitCount] + 1
		for i := 0; i <= carryDigitCount; i++ {
			cur[n-1-i] = min
		}
		return true, cur
	}
}

func computeScore(abcds []ABCD, seqA SeqA) int {
	totalD := 0
	for _, abcd := range abcds {
		if seqA[abcd.B-1]-seqA[abcd.A-1] == abcd.C {
			totalD += abcd.D
		}
	}
	return totalD
}

type SeqA []int

type ABCD struct {
	A int
	B int
	C int
	D int
}

// util
// *   math/simple

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func Min(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func Pow(base int, exponent uint) int {
	answer := 1
	for i := uint(0); i < exponent; i++ {
		answer *= base
	}
	return answer
}

// *   io

type Scanner struct {
	bufScanner *bufio.Scanner
}

func NewScanner() *Scanner {
	bufSc := bufio.NewScanner(os.Stdin)
	bufSc.Split(bufio.ScanWords)
	bufSc.Buffer(nil, 100000000)
	return &Scanner{bufSc}
}

func (sc *Scanner) ReadString() string {
	bufSc := sc.bufScanner
	bufSc.Scan()
	return bufSc.Text()
}

func (sc *Scanner) ReadInt() int {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}
