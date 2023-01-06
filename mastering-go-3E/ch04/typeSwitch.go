package main

import "fmt"

type Secret struct{
	pass string
}

type Entry struct {
	name string
	secret Secret
}

func matchType(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret Type")
	case Entry:
		fmt.Println("Entry Type")
	default:
		fmt.Printf("Not supported Type: %T\n",T)
	}
}

func learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("%v is %T type value\n",x,T)
	}
}

func main() {
	s := Secret{"passwd"}
	e := Entry{"rahul",s}

	matchType(s)
	matchType(e)
	learn(1)
	learn(1.0)
	learn("1")
	learn('1')
	matchType(1)
}