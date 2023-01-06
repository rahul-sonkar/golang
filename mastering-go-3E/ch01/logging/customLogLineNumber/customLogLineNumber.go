package main

import (
	"fmt"
	"os"
	"log"
	"path"
)

func main(){
	LOGFILE := path.Join(os.TempDir(),"myGo_customLineNumber")
	f,err := os.OpenFile(LOGFILE,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	myLog := log.New(f,"myLog ",log.Lshortfile)
	myLog.Println("Custome lines")
	myLog.Println("Custom line_numbers")
}
