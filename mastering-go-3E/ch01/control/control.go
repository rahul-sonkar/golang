package main

import (
	"fmt"
	"os"
	"strconv"

)
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provoide a argument.")
		return
	}
	argument := os.Args[1]

	switch argument{
	case "0":
		fmt.Println("Zeroo")
	case "1":
		fmt.Println("One")
	case "2","3","4":
		fmt.Println("Two, three or four")
		fallthrough
	default:
		fmt.Printf("Value is %s\n",argument)
	}

	value,err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Unable to convert string to int",err)
		return
	}
	
	switch {
	case value == 0:
		fmt.Println("Number is zero")
	case value > 0:
		fmt.Println("Number is positive")
	case value < 0:
		fmt.Println("Number is negtive")
	}
}