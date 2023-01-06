package main

import "fmt"

func main()  {
	// Creating an empty slice
	aSlice := []int{}
	// Printing info 
	fmt.Println("aSlice:",aSlice,"len:",len(aSlice),"cap:",cap(aSlice))

	// appending
	aSlice = append(aSlice, 1,2)
	// printing info 
	fmt.Println("aSlice:",aSlice,"len:",len(aSlice),"cap:",cap(aSlice))


	// A slice with length 4
	t := make([]int,4)
	t[0] = 1
	t[1] = 2
	t[2] = 3
	t[3] = 4

	// Need append() for adding more values because t's length is 4
	t = append(t, 5)
	fmt.Println("t:",t)

	// A 2D slice
	twoD := [][]int{{1,2,3},{4,5,6}}
	
	// Printing twoD values
	for _,i := range twoD {
		for _,r := range i {
			fmt.Println(r," ")
		}
		fmt.Println()
	}

	make2D := make([][]int,2)
	fmt.Println(make2D)
	make2D[0] = []int{1,2,3,4}
	make2D[1] = []int{5,6,7,8}
	fmt.Println(make2D)
 }