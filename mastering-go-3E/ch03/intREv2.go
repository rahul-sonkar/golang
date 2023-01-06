package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	byteSlice := []byte(s)
	re := regexp.MustCompile(`^\d+$`)
	return re.Match(byteSlice)
}

func main() {
	var (
		args = os.Args
		totalNumOfInput int
		totalNumOfTrue int
		totalNumOfFalse int
	)

	if len(args) == 1 {
		fmt.Println("Uses: ./intREv2.go arguments..")
		return
	}

	totalNumOfInput = len(args[1:])

	for _,v := range args[1:] {
		ok := matchInt(v)
		if ok {
			fmt.Println(v)
			totalNumOfTrue++
		} else {
			totalNumOfFalse++
		}
	}

	fmt.Println("Total number of input:",totalNumOfInput)
	fmt.Println("Total number of true input:",totalNumOfTrue)
	fmt.Println("Total number of false input:",totalNumOfFalse)
}