package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var (
	csvFile string = ""
	data    = []Entry{}
	index   map[string]int
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Uses: ./phonebook2.go  [ list|search|delet|insert ] argument")
		return
	}

	// csvFile = args[1]
	_,err := os.Stat(csvFile)
	if err != nil {
		fmt.Println("Creating",csvFile)
		f,err := os.Create(csvFile)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileInfo, err := os.Stat(csvFile)
	// Is it a regular file?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(csvFile,"is not a regular file.")
		return
	}
	
	err = readCsvFile(csvFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}


	// Differentiating between the commands.
	switch args[2] {
	case "insert":
		if len(args) != 5 {
			fmt.Println("Uses: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(args[4],"-","")

		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:",t)
			return
		}
		
		temp := initStruct(args[2],args[3],t)
		if temp != nil {
			err = insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		
	case "delete":
		if len(args) != 3 {
			fmt.Println("Uses: ./phonebook.go delete tel-number")
			return
		}

		t := strings.ReplaceAll(args[2],"-","")
		if !matchNum(t) {
			fmt.Println("Not a valid number.")
			return
		}

		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}

	case "search":
		if len(args) != 3 {
			fmt.Println("Uses: ./phonebook.go delete tel-number")
			return
		}

		t := strings.ReplaceAll(args[2],"-","")
		if !matchTel(t) {
			fmt.Println("Not a valid number.")
			return
		}

		temp := search(t)
		if temp == nil {
			fmt.Println("Number not found.")
			return
		}
		fmt.Println(*temp)
	
	case "list":
		list()
	
	default:
		fmt.Println("Not a valid option")
	}
}

func createIndex() error {
	index = make(map[string]int)
	for i,v := range data {
		key := v.Tel
		index[key] = i
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found.",key)
	}
	data = append(data[:i],data[i+1:]... )
	err := saveCsvFile(csvFile)
	if err != nil {
		return err
	}
	return nil
}


func insert(pS *Entry) error {
	// Check if it alredy exists.
	_,ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists.",pS.Tel)
	}
	data = append(data, *pS)
	//Update the index
	_ = createIndex()

	err := saveCsvFile(csvFile)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) *Entry {
	i,ok := index[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(),10)
	return &data[i]
}


func initStruct(n,s,t string) *Entry {
	if s == "" ||t == "" {
		return nil
	}

	lastAccess := strconv.FormatInt(time.Now().Unix(),10)
	return &Entry{Name: n,Surname: s,Tel: t,LastAccess: lastAccess}
}

func matchNum(t string) bool {
	b := []byte(t)
	re := regexp.MustCompile(`^\d+&`)
	return re.Match(b)
}

func readCsvFile(filepath string) error {
	_,err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f,err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines,err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Entry{
			Name: line[0],
			Surname: line[1],
			Tel: line[2],
			LastAccess: line[3],
		}
		data = append(data, temp)
	}

	return nil
}

func saveCsvFile(filepath string) error {
	csvfile,err := os.Create(filepath)
	if err != nil {
		return nil
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _,row := range data {
		temp := []string{row.Name,row.Surname,row.Tel,row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}


func list() {
	for _,v := range data {
		fmt.Println(v)
	}
}