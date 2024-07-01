package main

import (
	"fmt"
	"sort"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	if s == t {
		print("No")
		return
	}

	var ss []rune
	var is []int
	ss = []rune(s)
	for i := range ss {
		is = append(is, int(ss[i]))
	}

	var st []rune
	var it []int
	st = []rune(t)
	for i := range st {
		it = append(it, int(st[i]))
	}

	sort.Ints(is)
	sort.Ints(it)

	for i := range is {
		if i > len(it) {
			print("No")
			return
		}
		if is[i] > it[len(it)-i-1] {
			print("No")
			return
		} else if is[i] < it[len(it)-i-1] {
			print("Yes")
			return
		}
	}

	print("Yes")
}
