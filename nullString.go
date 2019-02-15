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
	DNS := user + ":" + pass + "@tcp(" + host + ":" + port + ")/testdb"
	//DNS := user + ":" + pass + "@tcp(" + host + ":" + port + ")/"
	fmt.Println("DNS:", DNS)
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	strQuery := "SELECT * FROM persons2"
	rows, err := db.Query(strQuery)
	if err != nil {
		log.Fatal("Error in Query, Error Info:", err)
	}
	for rows.Next() {
		var PersonID, LastName, FirstName, Address, City sql.NullString
		if err := rows.Scan(&PersonID, &LastName, &FirstName, &Address, &City); err != nil {
			fmt.Println("Error in Scan, Error Info:", err)
		}
		fmt.Println("PersonID:", PersonID, ";LastName:", LastName, ";FirstName:", FirstName, ";Address:", Address, ";City:", City)
		fmt.Println("PersonID.String:", PersonID.String, ";LastName.String:", LastName.String, ";FirstName.String:", FirstName.String, ";Address.String:", Address.String, ";City.String:", City.String)
	}
	defer rows.Close()

	fmt.Println("======================")
	var s sql.NullString
	strQuery = "SELECT LastName FROM persons2 WHERE PersonID=?"
	err = db.QueryRow(strQuery, 34).Scan(&s)
	if err != nil {
		log.Fatal("Error in Query, Error Info:", err)
	}
	if s.Valid{
		fmt.Println(">>>>>LastName:",s.String)
	}else {
		fmt.Println(">>>>>s.String is NULL:",s.String)
	}
}