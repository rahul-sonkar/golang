package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// Custom data type for storing details
type Details struct {
	Name       string
	SurName    string
	Tel        string
	LastAccess string
}

type entry []Details
var db entry
var index = make(map[string]int)
var csvFile = ""

// Implementing sort.Interface
func (e entry) Len() int {
	return len(e)
}

func (e entry) Less(i,j int) bool {
	if e[i].SurName == e[j].SurName {
		return e[i].Name < e[j].Name
	}
	return e[i].SurName < e[j].SurName
}

func (e entry) Swap(i,j int) {
	e[j],e[j] = e[j],e[i]
}

// Reading data from file.
func readCSVFile(filepath string) error {
	_,err := os.Stat(filepath)
	if err != nil {
		return err
	}

	openFile,err := os.Open(filepath)
	if err != nil {
		openFile.Close()
		return err
	}
	defer openFile.Close()

	csvReader,err := csv.NewReader(openFile).ReadAll()
	if len(csvReader) == 0 {
		return fmt.Errorf("%s :File is empty.",openFile)
	}
	if err != nil {
		return err
	}

	for _,line := range csvReader {
		temp := Details{
			Name: line[0],
			SurName: line[1],
			Tel: line[2],
			LastAccess: line[3],
		}
		db = append(db, temp)
	}
	return nil
}

// Saving data into the file localy.
func saveData(filepath string) error {
	file,err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for _,v := range db {
		lastAccess := strconv.FormatInt(time.Now().Unix(),10)
		temp := []string{v.Name,v.SurName,v.Tel,lastAccess}
		err = csvWriter.Write(temp)
		if err != nil {return err}
	}
	return nil
}

// Getting file from device.
func getFile() error {
	filePath := os.Getenv("PHONEBOOK")

	if filePath != "" {
		csvFile = filePath
		return nil
	}

	file,err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	return nil
}

// Creating database index
func creatIndex() {
	for i,v := range db {
		key := v.Tel
		index[key] = i
	}
}

// List: list all values of db
func list() {
	sort.Sort(db)
	for _,v := range db {
		fmt.Println(v)
	}
}

// Init Structure
func initStruct(n,s,t string) (*Details,error) {
	if n == "" || s == "" || t == "" {
		return nil,fmt.Errorf("felids are empty.")
	}
	temp := Details{
		Name: n,
		SurName: s,
		Tel: t,
		LastAccess: strconv.FormatInt(time.Now().Unix(),10),
	}
	return &temp,nil
}

// Insert records in db
func insert(n,s,t string) error {
	err := matchTel(t)
	if err != nil {return err}
	err = matchString(n)
	if err != nil {return err}
	err = matchString(s)
	if err != nil {return err}

	_,ok := index[t]
	if ok {
		return fmt.Errorf("Already exsit in database.")
	}

	temp,err := initStruct(n,s,t)
	if err != nil {
		return err
	}
	db = append(db,*temp)
	creatIndex()
	return nil
}

// Delete Data from database.
func delete(t string) error {
	err := matchTel(t)
	if err != nil {
		return err
	}
	i,ok := index[t]
	if !ok {
		return fmt.Errorf("details not found in database.")
	}
	db = append(db[:i],db[i+1:]...)
	return nil
}

// Search data
func serach(t string) error {
	err := matchTel(t)
	if err != nil {
		return err
	}
	i,ok := index[t]
	if !ok {
		return fmt.Errorf("details not found in database.")
	}
	fmt.Println(db[i])
	return nil
}

// Match the tel is valid telephone number or not.
func matchTel(t string) error {
	b := []byte(t)
	re := regexp.MustCompile(`^\d+$`)
	ok := re.Match(b)
	if !ok {
		return fmt.Errorf("Not valid telephone number.")
	}
	return nil
}

func matchString(n string) error {
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	ok := re.MatchString(n)
	if !ok {
		return fmt.Errorf("Not a valid name/surname.")
	}
	return nil
}


func init() {
	err := getFile()
	if err != nil {
		fmt.Println(err)
	}

	err = readCSVFile(csvFile)
	if err != nil {
		fmt.Println(err)
	}
}

func execute() {
	var err error
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Uses: ./phonebook.go list|search|delete|insert arguments..")
		return
	}

	switch args[1] {
	case "list":
		list()
	case "search":
		if len(args) != 3 {
			fmt.Println("Uses: search `telephone number`")
			return
		}
		err = serach(args[2])
		if err != nil {fmt.Println(err)}
	case "delete":
		if len(args) != 3 {
			fmt.Println("Uses: delete `telephone number`")
			return
		}
		err = delete(args[2])
		if err != nil {fmt.Println(err)}
	case "insert":
		if len(args) != 5 {
			fmt.Println("Uses: insert `Name` `Surname` `Telephone`")
			return
		}
		err = insert(args[2],args[3],args[4])
		if err != nil {
			fmt.Println(err)
		}
	}
}