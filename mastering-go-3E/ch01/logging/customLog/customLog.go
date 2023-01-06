package main

import (
	"fmt"
	"os"
	"log"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(),"myGo.log")
	f,err := os.OpenFile(LOGFILE,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f,"iLog ",log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 3rd Edition!")
}