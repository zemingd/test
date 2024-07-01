package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	L, R := ReadInt(), ReadInt()
	if R-L+1 >= 2019 {
		fmt.Println(0)
		return
	}
	if L%2019 < R%2019 {
		L %= 2019
		fmt.Println((L * (L + 1)) % 2019)
	} else {
		fmt.Println(0)
	}
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
