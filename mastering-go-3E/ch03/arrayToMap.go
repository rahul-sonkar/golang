package main

import "fmt"

func main() {
	array := [5]string{"a","B","c","D","e"}
	m := make(map[string]int)

	for i,v := range array {
		m[v] = i+1
	}
	fmt.Println(array)
	fmt.Println(m)
}