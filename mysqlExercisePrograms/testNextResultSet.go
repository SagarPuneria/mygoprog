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
	// multiStatement=true : allow multi statement in one query string.
	// interpolateParams=true : this is required to use placeholder in 2nd query.
	DNS := user + ":" + pass + "@tcp(" + host + ":" + port + ")/testdb?multiStatements=true&interpolateParams=true"

	/*DNS := user + ":" + pass + "@tcp(" + host + ":" + port + ")/testdb"
	db.Queru Error: 2019/02/02 15:10:33 Error in Query, Error Info:Error 1064:
	You have an error in your SQL syntax; check the manual that corresponds to your MySQL server
	version for the right syntax to use near 'SELECT * FROM persons' at line 1*/

	fmt.Println("DNS:", DNS)
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	strQuery := "SELECT 42; SELECT * FROM persons" // multi query which creates multi result set.
	rows, err := db.Query(strQuery)
	if err != nil {
		log.Fatal("Error in Query, Error Info:", err)
	}

	if !rows.NextResultSet() {
		log.Fatal("expected more result sets", rows.Err())
	}
	if rows.Next() {
		var PersonID, LastName, FirstName, Address, City string
		if err := rows.Scan(&PersonID, &LastName, &FirstName, &Address, &City); err != nil {
			fmt.Println("Error in Scan, Error Info:", err)
		}
		fmt.Println("PersonID:", PersonID, "LastName:", LastName, "FirstName:", FirstName, "Address:", Address, "City:", City)
	}
}
