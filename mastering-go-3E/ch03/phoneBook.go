// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"os"
// 	"regexp"
// 	"strings"
// 	"time"
// )

// type Entry struct {
// 	Name       string
// 	LastName   string
// 	Number     string
// 	LastAccess string
// }

// var csvfile = "./MOCK_DATA.csv"
// var data = []Entry{}
// var index = make(map[string]int)

// func loadData(filepath string) error {
// 	_, err := os.Stat(filepath)
// 	if err != nil {
// 		return err
// 	}

// 	open, err := os.Open(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	defer open.Close()

// 	file := csv.NewReader(open)
// 	for {
// 		slice, err := file.Read()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		temp := Entry{
// 			Name:       slice[0],
// 			LastName:   slice[1],
// 			Number:     slice[2],
// 			LastAccess: slice[3],
// 		}

// 		data = append(data, temp)
// 	}
// }

// func savefile() {
// 	file, err := os.Create(csvfile)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	csvfile := csv.NewWriter(file)
// 	for _, v := range data {
// 		temp := []string{v.Name, v.LastName, v.Number, v.LastAccess}
// 		_ = csvfile.Write(temp)
// 	}
// 	csvfile.Flush()
// }

// func creatIndex() {
// 	for i, v := range data {
// 		index[v.Number] = i
// 	}
// }

// func delete(num string) {
// 	n := index[num]
// 	fmt.Println("Deleting data..")
// 	data = append(data[:n], data[n+1:]...)
// 	fmt.Println("Data deleted...")
// 	fmt.Println("Recreating index...")
// 	creatIndex()
// 	fmt.Println("Index recreated...")
// 	savefile()
// }

// func insert(s string) {
// 	f := strings.Split(s, ",")
// 	var (
// 		name   string
// 		lname  string
// 		number string
// 	)

// 	if matchNames(f[0]) && matchNames(f[1]) {
// 		name = f[0]
// 		lname = f[1]
// 	}

// 	if matchNum(f[2]) {
// 		number = f[2]
// 	}

// 	temp := Entry{
// 		Name:       name,
// 		LastName:   lname,
// 		Number:     number,
// 		LastAccess: time.Now().Format("2022-04-05 06:47:59"),
// 	}

// 	data = append(data, temp)
// 	fmt.Println("Data inserted successfully...")
// 	fmt.Println("Recreating index...")
// 	creatIndex()
// 	fmt.Println("Index recreated...")
// 	savefile()

// }

// func matchNames(str string) bool {
// 	b := []byte(str)
// 	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
// 	return re.Match(b)
// }

// func matchNum(str string) bool {
// 	b := []byte(str)
// 	re := regexp.MustCompile(`^\d+$`)
// 	return re.Match(b)
// }

// func list() {
// 	for i, v := range data {
// 		fmt.Printf(
// 			"%d,  Name: %v, Surname: %v, Number: %v, LaseAccesse: %v\n",
// 			i+1, v.Name, v.LastName, v.Number, v.LastAccess)
// 	}
// }

// func search(s string) {
// 	n := index[s]
// 	if s == data[n].Number {
// 		fmt.Printf(
// 			"%d,  Name: %v, Surname: %v, Number: %v, LaseAccesse: %v\n",
// 			n+1, data[n].Name, data[n].LastName, data[n].Number, data[n].LastAccess)
// 	} else {
// 		fmt.Println("Not found.")
// 	}
// }

// func init() {
// 	err := loadData(csvfile)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Data loaded in database...")
// 	fmt.Println("Creating index ....")
// 	creatIndex()
// 	fmt.Println("Index created...")
// }
// func main() {
// 	args := os.Args
// 	if len(args) == 1 {
// 		fmt.Println("Uses: ./phoneBook.go list|search|delete|insert  argument")
// 		return
// 	}

// 	switch args[1] {
// 	case "list":
// 		list()
// 	case "search":
// 		if len(args) != 3 {
// 			fmt.Println("Required Number.")
// 			return
// 		}
// 		search(args[2])
// 	case "delete":
// 		if len(args) != 3 {
// 			fmt.Println("Required Number.")
// 			return
// 		}
// 		delete(args[2])
// 	case "insert":
// 		if len(args) != 3 {
// 			fmt.Println("Required data.")
// 			return
// 		}
// 		n := index[args[2]]
// 		if args[2] == data[n].Number {
// 			fmt.Println(args[2], "already exists.")
// 			return
// 		}
// 		insert(args[2])
// 	default:
// 		fmt.Println("Invalid options..")
// 		return
// 	}
// }
