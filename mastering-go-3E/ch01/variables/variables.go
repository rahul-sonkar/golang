package main

import "fmt"

// Global variabel
var global int = 1

// Global constant
const G int = 100

func main() {
	//local variable using `var` key word without intialize
	var a int
	
	//local variable using `var` key word with intialize
	var b int = 1

	//using short assignment sentext
	c := 2

	fmt.Println("a",a,"b",b,"c",c,"global var",global,"global constant",G)
}