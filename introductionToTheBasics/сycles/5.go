package main

import "fmt"

func main(){

    var n,c,d int
    fmt.Scan(&n,&c,&d)
    if (n % c == 0) && (n % d != 0){
        fmt.Println(n)
    }else if (c % c == 0) && (c % d != 0){
        fmt.Println(c)
    }else if (d % c == 0) && (d % d != 0){
        fmt.Println(d)
    }

}