package main

import "fmt"

func main(){
    var a int
    a = 123
    fmt.Scan(&a)
    fmt.Println(a % 10)
}
