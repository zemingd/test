package main

import (
    "fmt"
    "bufio"
    "os"
)
var r = bufio.NewReader(os.Stdin)

func main(){
    var a, b, c int
    fmt.Fscan(r, &a, &b, &c)
    if a < b && b < c {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}

