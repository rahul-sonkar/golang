package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name string
	Surname string
	Number string
	LastAccess string
}

var myData = []Record{}

func readCSVFile(filepath string) ([][]string,error){
	_,err := os.Stat(filepath)
	if err != nil {
		return nil,err
	}

	f,err := os.Open(filepath)
	if err != nil {
		return nil,err
	}
	defer f.Close()

	lines,err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil,err
	}

	return lines,err
}

func saveCSVFile(filepath string) error {
	csvfile,err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	// Changing the default field delimeter to tab
	// csvwriter.Comma = '\t'
	for _,row := range myData {
		temp := []string{row.Name,row.Surname,row.Number,row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func main()  {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Uses: csvData input output")
		return
	}

	input := args[1]
	output := args[2]

	lines,err := readCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _,line := range lines {
		temp := Record{
			Name: line[0],
			Surname: line[1],
			Number: line[2],
			LastAccess: line[3],
		}

		myData = append(myData, temp)
		fmt.Println(temp)
	}

	err = saveCSVFile(output)
	if err != nil {
		fmt.Println(err)
		return
	} 
}