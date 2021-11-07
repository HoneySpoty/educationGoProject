package main

import "fmt"

func main(){
var n int
    fmt.Scan(&n)
    var s int
    for i:=0; i<n; i++{
        var t int
        fmt.Scan(&t)
        if t>=16 && t<=96 && t%8==0 {
            s+=t
        }
    }
    fmt.Print(s)
}