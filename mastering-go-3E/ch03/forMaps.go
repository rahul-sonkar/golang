package main

import "fmt"

func main() {
	aMap := make(map[string]int)
	aMap["key1"] = 1
	aMap["key2"] = 2
	aMap["key3"] = 3

	for key, value := range aMap {
		fmt.Println(key, ":", value)
	}
}
