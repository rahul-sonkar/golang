package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo,err := os.Stat("./test.sh")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = fileInfo.Mode()
	
}
