package main

import "fmt"


type s1 struct {
	name string
	age int
}

type s2 struct {
	tel string
}

func print(a interface{}) {
	fmt.Println(a)
}

func main() {
	a := s1{"rahul",20}
	b := s2{"9131535028"}

	print(a)
	print(b)
	print(123)
	print("Go programming language..")
}