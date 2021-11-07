package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := a / 1000
	c := a % 1000
	bsum := (b % 10) + (b / 100) + (b / 10 % 10)
	csum := (c % 10) + (c / 100) + (c / 10 % 10)
	if(bsum == csum){
		fmt.Print("YES")
	}else {
		fmt.Print("NO")
	}
}