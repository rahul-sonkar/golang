package main

import (
	"fmt"
	"os"
)

type info struct {
	name    string
	surname string
	tel     int
}

var phone_book []info

func list() {
	for _, v := range phone_book {
		fmt.Println(v)
	}
}

func search(s string) {
	for _, v := range phone_book {
		if s == v.surname {
			fmt.Println(v)
			return
		}
	}
	fmt.Println("Not found.", s)
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Uses: ./phonebook.go list|[search <surname>]")
		return
	}

	phone_book = append(phone_book, info{"rahul", "sonkar", 001})
	phone_book = append(phone_book, info{"giga", "sen", 002})
	phone_book = append(phone_book, info{"packet", "roy", 003})

	switch args[1] {
	case "list":
		list()
	case "search":
		if len(args) != 3 {
			fmt.Printf("Uses: search <surname>\n")
			return
		}
		search(args[2])
	default:
		fmt.Println("Invalid options.")
	}
}
