package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Required a record.")
		return
	}

	b := matchRecord(args[1])

	if b {
		fmt.Println(args[1],"is correct")
	} else {
		fmt.Println(args[1],"isn't correct")
	}
}

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func matchTel(i string) bool {
	t := []byte(i)
	re := regexp.MustCompile(`^[0-9]*$`)
	return re.Match(t)
}

func matchRecord(s string) bool {
	fields := strings.Split(s,",")
	if len(fields) != 3{
		return false
	}

	if !matchNameSur(fields[0]) {
		return false
	}

	if !matchNameSur(fields[1]) {
		return false
	}
	 return matchTel(fields[2])
}