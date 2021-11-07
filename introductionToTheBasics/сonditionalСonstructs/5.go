package main

import "fmt"

func main(){
    var n rune
    fmt.Scan(&n)
    
    switch{
        case n%400==0 || n%4==0 && n%100!=0: fmt.Print("YES")
        default: fmt.Print("NO")
    }
}