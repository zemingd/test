package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := ReadInt()
	A := ReadInts(N)
	B := ReadInts(N)
	C := ReadInts(N - 1)
	prev := -2
	ans := 0
	for i := 0; i < N; i++ {
		a := A[i] - 1
		ans += B[a]
		if prev == a-1 {
			ans += C[prev]
		}
		prev = a
	}
	fmt.Println(ans)
}

var reader = bufio.NewReader(os.Stdin)

func Scan(a ...interface{}) {
	if _, err := fmt.Fscan(reader, a...); err != nil {
		panic(err)
	}
}
func ReadInt() (i int)       { Scan(&i); return }
func ReadString() (s string) { Scan(&s); return }
func ReadInts(n int) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		Scan(&v[i])
	}
	return v
}
