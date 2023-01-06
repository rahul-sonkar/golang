package main

import "fmt"

func fn() (int,int) {
	return 1,2
}

func main() {
	a,b := fn()
	fmt.Printf("%d %d\n",a,b)
}