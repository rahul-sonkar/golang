package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	_ "github.com/lib/pq"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname port username password db")
	return
	}

	host := arguments[1]
	p := arguments[2]
	user := arguments[3]
	pass := arguments[4]
	database := arguments[5]

	// Port number should be an integer
	port,err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Not a valid port number:",err)
		return
	}

	// connection string
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, database)

	// open PostgreSQL database
	db, err := sql.Open("postgres",conn)
	if err != nil {
		fmt.Println("open()",err)
		return
	}
	defer db.Close()

	// Get all databases
	row,err := db.Query(`SELECT "datname" FROM "pg_database" WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Query",err)
		return
	}
	
	//
}