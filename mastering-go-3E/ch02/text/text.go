package main

import "fmt"

func main() {
	aStr := "Hello world!"
	fmt.Println("First chr:",string(aStr[0]))

	chr := '$'
	fmt.Println("As an int32 value:",chr)
	fmt.Printf("chr :%c\n",chr)

	for _,v := range aStr {
		fmt.Printf("%x ",v)
	}
	fmt.Println()

	for _,v := range aStr {
		fmt.Printf("%c",v)
	}
	fmt.Println()
}