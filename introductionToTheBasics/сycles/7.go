package main

import "fmt"

func main() {
	var vklad,procent,maxv int
	fmt.Scan (&vklad,&procent, &maxv)
	for god := 1; vklad < maxv; god++{
		vklad = vklad + (procent*vklad)/100
			if vklad > maxv {
			fmt.Println(god)
			break
		}
	}
}