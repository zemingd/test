package main

import "fmt"
import "strings"
import "strconv"

func swap(a, b int)(int, int){
    return b, a 
}

func main(){

    var s []string
    var n []int
    var in_s string

    fmt.Scanln(&in_s)

    s = strings.Fields(in_s)
    
    for i, v := range s{
        n = append(a, strconv.Atoi(v))
    }
 
    for i:= 0; i < len(a) - 1; i++ {
        for j := i + 1 ; j < len(a) ; j++{
            n[i], n[j] = swap(n[i], n[j])
        }
    } 
    
    fmt.Println(n[0], n[1], n[2])


}
