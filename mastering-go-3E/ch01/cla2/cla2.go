package main
import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provoid arguments.")
		return
	}
	var min, max int
	for i,v := range args {
		value,_ := strconv.Atoi(v)
		if i == 1 {
			min = value
			max = value
			continue
		}
		
		if min > value {
			min = value
			continue
		}
		if max < value {
			max = value
			continue
		}
	}
	fmt.Printf("min_value is: %d\nmax_value is: %d\n",min,max)
}