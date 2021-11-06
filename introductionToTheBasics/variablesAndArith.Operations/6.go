package main

import "fmt"

func main() {
    var d int
    fmt.Scan(&d)
    minutes := d * 12 * 60 / 360
    h := minutes / 60
    m := minutes % 60
    fmt.Println("It is", h, "hours", m, "minutes.")
}    