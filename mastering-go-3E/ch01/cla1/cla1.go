package main
import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a argument.")
		return
	}

	arg := os.Args[1]
	value,err := strconv.ParseInt(arg,10,64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Your value is:",value)
}