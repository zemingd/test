package main

import (
	"fmt"
	"sync"
)

type graph struct {
	path []int
	height int
	// flag bool // pathに存在するもので自分より高いものがあればtrue
}
var graphs []graph
func main() {
	var N,M int
	fmt.Scan(&N,&M)
	graphs = make([]graph,N)
	for i:= 0; i<N; i++ {
		var h int
		fmt.Scan(&h)
		graphs[i].height = h
		graphs[i].path = []int{}
		// graphs[i].flag = false
	}
	for i:= 0; i<M; i++ {
		var a,b int
		fmt.Scan(&a,&b)
		graphs[a-1].path = append(graphs[a-1].path,b-1)
		graphs[b-1].path = append(graphs[b-1].path,a-1)
	}
	count := 0
	var wg sync.WaitGroup
	for i:=0; i<N; i++ {
		wg.Add(1)
		go func(i int) {
			paths := graphs[i].path
			height := graphs[i].height
			bestHeight := true // 一番高いのならtrue
			for _, value := range paths {
				if height <= graphs[value].height {
					bestHeight = false
					break
				}
			}
			if bestHeight {
				// fmt.Println(graphs[i],"ベスト")
				count++
			}
			defer wg.Done()
		}(i)
	}
	fmt.Println(count)
}
