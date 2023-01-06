package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) == 1 {
		fmt.Println("Need date..")
		return
	}

	dateStr := os.Args[1]

	// is this date only ?
	d, err := time.Parse("02 January 2006", dateStr)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

	// is this date + time ?
	d, err = time.Parse("02 January 2006 15:04", dateStr)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("date:", d.Day(), d.Month(), d.Year())
		fmt.Println("time:", d.Hour(), d.Minute())
	}

	// Is this a date + time with month represented as a number?
	d, err = time.Parse("02-01-2006 15:06", dateStr)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Is it time only?
	d, err = time.Parse("15:04", dateStr)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)

	// Convert Epoch time to time.Time
	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())

	duration := time.Since(start)
	fmt.Println("Execution time:",duration)
}
