package main

import (
	"fmt"
	"regexp"
)

func main()  {
	t := []byte("rahul")
	re := regexp.MustCompile("^[A-Z][a-z]*$")
	b := re.Match(t)
	fmt.Println(b)
}