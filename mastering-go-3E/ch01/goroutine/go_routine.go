package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go print()

	}
	time.Sleep(1 * time.Millisecond)
}
func print() {
	fmt.Println("printing from go routine.")
}
