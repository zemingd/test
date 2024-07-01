package main
 
import (
  "fmt"
  "strconv"
)

func pow (n int, r int) int {
  
  x := 1
  for i := 0; i < r; i++ {
    x *= n
  }
  return x
}

func main() {
  
  //入力値を分割
  var str string
  fmt.Scan(&str)
  var m [4]int
  for i := 0; i < 4; i++ {
    m[i],_ = strconv.Atoi(string(str[i]))
  }
  
  //+,-を格納
  var s [3]string
  
  //全パターン試して正解を見つけたら終了
  for i := 0; i < pow(2,3); i++ {
    ans := m[0]
    for j := 0; j < 3; j++ {
    //p := i & (1<<(uint(j))) 
      p :=  (i>>(uint(j))) & 1 
      if p == 0 {
        s[j] = "-"
        ans -= m[j+1]
      } else {
        s[j] = "+"
        ans += m[j+1]
      }
    }
    if ans == 7 { 
      fmt.Printf("%d%s%d%s%d%s%d=7", m[0], s[0], m[1], s[1], m[2], s[2], m[3])
      break
    }
  }
}