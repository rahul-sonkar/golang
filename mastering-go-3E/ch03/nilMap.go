package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test1"] = 1
	fmt.Println(aMap)

	aMap = nil
	fmt.Println(aMap)

	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
		aMap["from if block"] = 2
		fmt.Println(aMap)
	}

	// Give an error, program carsh
	aMap = nil 
	aMap["test2"] = 3 

}
