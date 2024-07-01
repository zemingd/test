package main

import (
        "bufio"
        "fmt"
        "os"
)

var rdr = bufio.NewReader(os.Stdin)

func readLine() string {
        var buf []byte
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
        return string(buf)
}

func lowerBound(a []int, x int) int {
        n := len(a)
        ok := n
        ng := -1
        for ok-ng > 1 {
                mid := (ok + ng) / 2
                if a[mid] >= x {
                        ok = mid
                } else {
                        ng = mid
                }
        }
        return ok
}

func main() {
        var s string
        var t string

        s = readLine()
        t = readLine()

        idxs := make([][]int, 256)

        for i := 0; i < len(s); i++ {
                idxs[s[i]] = append(idxs[s[i]], i)
        }

        n := int64(len(s))
        cur := int64(0)
        ng := false

        for i := 0; i < len(t); i++ {
                is := idxs[t[i]]
                if is == nil || len(is) == 0 {
                        ng = true
                        break
                }
                base := (cur / n) * n

                nidx := lowerBound(is, int(cur%n))
                if nidx == len(is) {
                        cur = base + int64(is[0]) + n + 1
                } else {
                        cur = base + int64(is[nidx]) + 1
                }
        }
        if ng {
                fmt.Println(-1)
        } else {
                fmt.Println(cur)
        }
}
