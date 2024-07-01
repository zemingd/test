package main

import (
	"bufio"
	"fmt"
	"os"
)

var H, W int

func main() {
	fmt.Scan(&H, &W)

	reader := bufio.NewScanner(os.Stdin)
	q := make([][]int, 0)
	m := make([][]bool, H)

	for i := 0; i < H; i++ {
		m[i] = make([]bool, W) // Falseで初期化されている
		reader.Scan()

		for j := 0; j < W; j++ {
			if string(reader.Text()[j]) == "#" {
				q = append(q, []int{i, j})
				m[i][j] = true
			}
		}
	}
	fmt.Println(bfs(q, m))
}

func bfs(q [][]int, m [][]bool) int {
	var ans int

	for len(q) != 0 {
		ans++
		len := len(q)

		for i := 0; i < len; i++ {
			var y, x int = q[0][0], q[0][1]
			q = append(q[1:])

			if 0 <= y+1 && y+1 < H && m[y+1][x] == false {
				m[y+1][x] = true
				q = append(q, []int{y + 1, x})
			}
			if 0 <= y-1 && y-1 < H && m[y-1][x] == false {
				m[y-1][x] = true
				q = append(q, []int{y - 1, x})
			}
			if 0 <= x+1 && x+1 < W && m[y][x+1] == false {
				m[y][x+1] = true
				q = append(q, []int{y, x + 1})
			}
			if 0 <= x-1 && x-1 < W && m[y][x-1] == false {
				m[y][x-1] = true
				q = append(q, []int{y, x - 1})
			}
		}
	}
	return ans - 1
}
