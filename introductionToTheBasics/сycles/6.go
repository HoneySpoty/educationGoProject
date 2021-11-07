package main

import "fmt"

func main(){
    for i:=0; i<=100; fmt.Scan(&i) {
        if i>=10 {
            fmt.Println(i)
        }
    }
