package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer fmt.Print(1)
	var nm = readInts()
	n := nm[0]
	m := nm[1]

	var storea = make([]int, m, m)
	var storeb = make([]int, m, m)
	for i := 0; i < m; i++ {
		var ab = readInts()
		storea[i] = ab[0]
		storeb[i] = ab[1]
	}

	var connection = make([][]int, n, n)
	for i := 0; i < m; i++ {
		a := storea[i] - 1
		b := storea[i] - 1
		connection[a] = append(connection[a], b)
		connection[b] = append(connection[b], a)
	}

	var groupIDs = make([]int, n)
	for i := 0; i < n; i++ {
		groupIDs[i] = -1
	}

	var marking func(i int, group int)
	marking = func(i int, group int) {
		if groupIDs[i] != -1 {
			return
		}
		groupIDs[i] = group
		for _, next := range connection[i] {
			marking(next, group)
		}
	}
	for i := 0; i < n; i++ {
		marking(i, i)
	}
	var counter = make([]int, n, n)
	for _, v := range groupIDs {
		counter[v]++
	}
	var temp = 0
	for _, v := range counter {
		if temp < v {
			temp = v
		}
	}
	fmt.Printf("%v\n", temp)
}

func readLine() []string {
	rdr := bufio.NewReaderSize(os.Stdin, 1000)
	buf := make([]byte, 1000, 1000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return strings.Split(string(buf), " ")
}

func readInts() []int {
	var strs = readLine()
	var ret []int
	for _, v := range strs {
		c, _ := strconv.Atoi(v)
		ret = append(ret, c)
	}
	return ret
}
