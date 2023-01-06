package main

import (
	"fmt"
	"regexp"
)

func main()  {
	t := []byte("123")	
	RE := regexp.MustCompile(`^[-+]?\d+$`)
	b := RE.Match(t)
	fmt.Println(b)
}