package main

import "fmt"

func print[T any](s []T) {
	for _,v := range s {
		fmt.Println(v)
	}
	fmt.Println()
}

func main() {
	ints := []int{1,2,3,4,5}
	strings := []string{"a","b","c","d"}
	print(ints)
	print(strings)
}