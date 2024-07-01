package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	defer out.Flush()

	s := next()
	t := next()

	sSlice := strings.Split(s, "")
	tSlice := strings.Split(t, "")

	min := len(tSlice)

	for i := 0; i < len(sSlice)-2; i++ {
		size := len(tSlice)
		if t[0] == s[i] {
			size--
		}
		if t[1] == s[i+1] {
			size--
		}
		if t[2] == s[i+2] {
			size--
		}
		if min > size {
			min = size
		}
	}
	fmt.Fprintf(out, "%d\n", min)
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}

func nextFloat64() float64 {
	a, _ := strconv.ParseFloat(next(), 64)
	return a
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func quickSort(val []int) []int {
	if len(val) < 2 {
		return val
	}

	pivot := val[0]

	left := []int{}
	right := []int{}

	for _, v := range val[1:] {
		if pivot > v {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	ret := append(left, pivot)
	ret = append(ret, right...)

	return ret
}
