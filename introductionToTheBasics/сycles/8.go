package main

import "fmt"

func main() {
  var a,b,rev int
  fmt.Scan(&a,&b)
  for a > 0 {
    rev = rev*10 + a%10
    a /=10
  }
  for i := rev; i > 0; i /= 10 {
    for j := b; j > 0; j /= 10 {
      if i%10 == j%10 {
        fmt.Print(i%10, " ")
      }
    }
  }
}