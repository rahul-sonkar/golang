package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type person struct {
	fname string
	lname string
	tel   int
}

var data = make([]person, 0)

// var data []person

func main() {
	arg := os.Args
	if len(arg) == 1 {
		fmt.Println("Uses: ./phonebook.go options{search <tel-number>|list}")
		return
	}

	fillData()

	switch arg[1] {
	case "search":
		if len(arg) != 3 {
			fmt.Println("Uses: ./phonebook.go search <tel-number>")
			return
		}
		num, err := strconv.Atoi(arg[2])
		if err != nil {
			fmt.Println("Please check telephone number once :D")
			return
		}
		person, err := search(num)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*person)

	case "list":
		list()

	default:
		fmt.Println("Invalid options.")
		fmt.Println("Uses: ./phonebook.go options{search <tel-number>|list}")
	}

}
func search(n int) (*person, error) {
	for _, v := range data {
		if v.tel == n {
			return &v, nil
		}
	}
	return &person{}, fmt.Errorf("not found in database")
}

func list() {
	for i, v := range data {
		fmt.Println(i,v)
	}
}
func fillData() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	for i := 0; i < 100; i++ {

		fname := make([]byte, 0)
		lname := make([]byte, 0)
		for j := 0; j < 4; j++ {
			fname = append(fname, byte((rand.Intn(91-65))+65))
		}

		for k := 0; k < 5; k++ {
			lname = append(lname, byte((rand.Intn(91-65))+65))
		}

		num := rand.Intn(201-100) + 100

		data = append(data, person{fname: string(fname), lname: string(lname), tel: num})
	}
}
