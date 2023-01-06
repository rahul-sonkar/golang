package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need more argument.")
		return
	}

	var total, nInt, nFloat int
	invalid := make([]string, 0)

	for _, k := range args[1:] {

		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInt++
			continue
		}

		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloat++
			continue
		}

		invalid = append(invalid, k)
	}
	fmt.Printf("read: %d   #nInt: %d   #nFloat: %d\n",total,nInt,nFloat)
	if len(invalid) != total {
		fmt.Println("Unknow arguments are:",len(invalid))
		for _,s := range invalid {
			fmt.Println(s)
		}
	}
}