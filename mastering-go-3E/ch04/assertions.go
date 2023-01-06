package main

import "fmt"

func returnInt() interface{} {
	return 12
}

func main() {
	anInt := returnInt()

	num := anInt.(int)
	num++
	fmt.Println(num)

	n,ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successfully",n)
	} else {
		fmt.Println("Type assertion unsuccessful.")
	}

	// Give an panic
	_ = anInt.(string)
}