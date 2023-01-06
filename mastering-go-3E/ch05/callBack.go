package main

import (
	"fmt"
	"net/http"
)

var m = make(map[string]func())

func add() {
	fmt.Println("I'm add()")
}

func sub() {
	fmt.Println("I'm sub()")
}

func callBack(note string,fn func()) {
	fmt.Println(note)
	fn()
}

func main() {
	m["add"] = add
	m["sub"] = sub

	callBack("I'm learning Go nowdays.",func(){
		res,err := http.Get("https://google.com")
		if err != nil {fmt.Println(err)}
		fmt.Println(res.Status)
	})
}
