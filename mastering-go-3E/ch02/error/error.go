package main

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)

func check(a,b int) error {
	if a == 0 && b == 0 {
		return errors.New("custom error msg")
	}
	return nil
}

func formattedError(a,b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("formatted error msg: a = %v, b = %v, UserId: %v",a,b,os.Getuid())
	}
	return nil
}

func main() {
	err := check(1,0)
	if err == nil {
		fmt.Println("check() is Ok")
	} else {
	fmt.Println(err)
	}
	err = formattedError(0,0)
	if err != nil {
		fmt.Println(err)
	}

	i,err := strconv.Atoi("-12")
	if err == nil {
		fmt.Println(i)
	} else {
	fmt.Println(err)
	}

	i,err = strconv.Atoi("y12")
	if err == nil {
		fmt.Println(i)
	} else {
	fmt.Println(err)
	}
}