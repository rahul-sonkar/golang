package main

import "fmt"

func main() {
	m :=make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	var (
		a1 []string
		a2 []int
	)

	for k, v := range m {
		a1 = append(a1, k)
		a2 = append(a2, v)
	}
	fmt.Println("map",m)
	fmt.Println("slice1",a1)
	fmt.Println("slice2",a2)
}
