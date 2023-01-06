package main

import (
	"fmt"
	"os"
)

type S struct {
	index int
	argument string
}

var args []S

func main() {
	for i,v := range os.Args {
		temp := S{
			index: i,
			argument: v,
		}
		args = append(args,temp)
	}
	fmt.Println(args)
}