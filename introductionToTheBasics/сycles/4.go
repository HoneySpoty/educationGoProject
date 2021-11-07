package main

import "fmt"

func main(){
    var a, max, count int  
   
    for fmt.Scan(&a); a != 0 ; fmt.Scan(&a){ 
        switch {
            case a > max:
            max, count = a, 1
            case a == max:
            count++
        }
    }
    fmt.Println(count)
}