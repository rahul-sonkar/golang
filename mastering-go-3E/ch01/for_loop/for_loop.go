package main
import "fmt"
func main() {
	// traditional for loop 
	for i:=0; i<10; i++ {
		fmt.Print(i*i," ")
	}
	fmt.Println()

	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i," ")
		i++
	}
	fmt.Println()

	// for loop used as while loop
	i = 0
	for {
		if i == 10 {break}
		fmt.Print(i*i," ")
		i++
	}
	fmt.Println()

	// for loop with range key word.
	aSlice := []int{1,5,6,7,8}
	for i,v := range aSlice {
		fmt.Println("index",i," ","value:",v)
	}
	fmt.Println()
}