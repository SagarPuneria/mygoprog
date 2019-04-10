package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := os.Args[1]
	pass := os.Args[2]
	host := os.Args[3]
	port := os.Args[4]
	DNS := user + ":" + pass + "@tcp(" + host + ":" + port + ")/testdb?multiStatements=true&interpolateParams=true"
	fmt.Println("DNS:", DNS)
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	strQuery := "SELECT * FROM persons;SELECT * FROM persons3;"
	rows, err := db.Query(strQuery)
	if err != nil {
		log.Fatal("Error in Query, Error Info:", err)
	}
	for rows.Next() {
		var PersonID, LastName, FirstName, Address, City string
		if err := rows.Scan(&PersonID, &LastName, &FirstName, &Address, &City); err != nil {
			fmt.Println("Error in Scan, Error Info:", err)
		}
		fmt.Println("PersonID:", PersonID, "LastName:", LastName, "FirstName:", FirstName, "Address:", Address, "City:", City)
	}
	if !rows.NextResultSet() { //when you have multi result set, you need to call NextResultSet() once to get 2nd result set.
		log.Fatal("expected more result sets", rows.Err())
	}
	for rows.Next() {
		var PersonID, LastName, FirstName, Address, City string
		if err := rows.Scan(&PersonID, &LastName, &FirstName, &Address, &City); err != nil {
			fmt.Println("Error in Scan, Error Info:", err)
		}
		fmt.Println("PersonID:", PersonID, "LastName:", LastName, "FirstName:", FirstName, "Address:", Address, "City:", City)
	}
}
