package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func contains(target int, list []int) bool {
	for _, r := range list {
		if target == r {
			return true
		}
	}
	return false
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()

	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func nextInt64() int64 {
	sc.Scan()

	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func nextF64() float64 {
	sc.Scan()

	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func isInteger(x float64) bool {
	return math.Floor(x) == x
}

func Ave(s []int) int {
	val := 0
	for _, i := range s {
		val += i
	}
	return int(val / len(s))
}
func cal(i int) (sum int) {
	for i != 0 {
		sum += i % 10
		i /= 10
	}
	return
}
func max(a []int) (idx int, val int) {
	max := a[0]
	idx = 0
	for j, i := range a {
		if max < i {
			max = i
			idx = j
		}
	}
	return idx, max
}

func draw(i int, a []int) (l []int) {
	// p = a[i]
	for j := i; j < len(a); j++ {
		if j != len(a)-1 {
			a[j] = a[j+1]
		} else if j == len(a)-1 {
			a = a[:len(a)-1]
		}
	}
	return a
}
func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

func f(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func main() {
	sc.Split(bufio.ScanWords)
	// sc.Buffer([]byte{}, math.MaxInt64)
	// a := make([][]int, 3)
	// for i := 0; i < 3; i++ {
	// 	a[i] = make([]int, 3)
	// }
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		a[i][j] = nextInt()
	// 	}
	// }

	s := nextInt()
	var a []int

	for i := 0; true; i++ {
		if i == 0 {
			a = append(a, s)
		} else if i > 0 {
			if contains(f(a[i-1]), a) {
				fmt.Println(i + 1)
				break
			}
			a = append(a, f(a[i-1]))
		}
	}

}
