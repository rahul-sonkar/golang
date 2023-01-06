package main

import (
	"fmt"
	"os"
	"regexp"
)

func checkNameSur(s string) bool {
	byteSlice := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(byteSlice)
}

func main() {
	args := os.Args
	if len(args) == 1{
		fmt.Println("Uses: ./nameSurREv2.go arguments..")
		return
	}

	for _,v := range args[1:] {
		ok := checkNameSur(v)
		if ok {
			fmt.Println(v,"is valid string.")
		} else {
			fmt.Println(v,"isn't valid string.")
		}
	}
}