package main

import (
	"fmt"
	"math"
)

func main() {
	
	radius := 3.4
	a := area(radius)
	c := circumferenc(radius)
	fmt.Printf("Area: %.2f\nCircumferenc: %.2f\n",a,c)

}

func area(r float64) float64 {
	area := math.Pi*(r*r)
	return area
}

func circumferenc(r float64) float64 {
	c := 2*math.Pi*r
	return c
}