package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt() //挿入回数
	K := nextInt() //出力する値の位置
	values := make([]int, N)
	counts := make(map[int]int, N)
	for i := 0; i < N; i++ {
		values[i] = nextInt()
		counts[values[i]] += nextInt()
	}
	sort.Ints(values)
	var current int
	prev_value := 0
	for _, v := range values {
		if prev_value == v {
			continue
		}
		current += counts[v]
		if K <= current {
			fmt.Println(v)
			break
		}
		prev_value = v
	}
}

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}
