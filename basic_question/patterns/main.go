package main

import (
	"fmt"
)

func main() {
	num := 7
	p3(num)
}

func p3(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			if j==i {
				fmt.Printf("\\")
			} else if j == (num-i+1) {
				fmt.Printf("/")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}
}

func p2(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i==0 || i==(num-1) || j==0 || j==(num-1) {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
 		}
		fmt.Println()
	}
}
func p1(num int) {

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
