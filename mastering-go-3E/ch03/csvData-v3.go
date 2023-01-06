package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type info struct {
	Name string
	Tel string
}

var store []info


func read(filepath string) error {
	_,err := os.Stat(filepath)
	if err != nil {
		return err
	}
	file,err := os.Open(filepath)
	if err != nil {
		file.Close()
		return err
	}
	defer file.Close()
	lines,err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}
	for _,line := range lines {
		temp := info{
			Name: line[0],
			Tel: line[1],
		}
		store = append(store, temp)
	}
	return nil
}

func write(filepath string,c rune) error {
	file,err := os.Create(filepath)
	if err != nil {
		file.Close()
		return err
	}
	defer file.Close()
	csvfile := csv.NewWriter(file)
	csvfile.Comma = rune(c)
	for _,v := range store {
		temp := []string{v.Name,v.Tel}
		err = csvfile.Write(temp)
		if err != nil {
			return err
		}
	}
	csvfile.Flush()
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1{
		fmt.Println("Uses: ./csvData-v2.go arguments..")
		return
	}
	chr := []rune(args[1])
	inputPath := args[2]
	outputPath := args[3]

	fileInfo,err := os.Stat(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(inputPath,"is not regular file.")
		return
	}

	err = read(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = write(outputPath,chr[0])
	if err != nil {
		fmt.Println(err)
		return
	}
}