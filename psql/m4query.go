package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host       = "127.0.0.1"
	port       = 5432
	user       = "postgres"
	password   = "postgres"
	dbName     = "postgres"
	schemaName = "employee"
	sslmode    = "disable"
)

func main() {

	hostPtr := flag.String("host", host, "a host flag")
	dbNamePtr := flag.String("dbName", dbName, "a dbName flag")
	schemaNamePtr := flag.String("schemaName", schemaName, "a schemaName flag")
	userPtr := flag.String("user", user, "a user flag")
	passwordPtr := flag.String("password", password, "a password flag")
	sslmodePtr := flag.String("sslmode", sslmode, "a user flag")
	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		*hostPtr, port, *userPtr, *passwordPtr, *dbNamePtr, *sslmodePtr)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	/*strQuery := `CREATE TABLE employee.Persons2 (PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255));
	INSERT INTO Persons2 (PersonID) VALUES (?);`
	rows, err := db.Query(strQuery, 32)*/

	/*strQuery := fmt.Sprintf("CREATE TABLE %s.Persons2 (PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255))", *schemaNamePtr)
	rows, err := db.Query(strQuery)*/

	/*strQuery := fmt.Sprintf("SELECT * FROM %s.Persons2 WHERE PersonID=? OR PersonID=?", *schemaNamePtr)
	rows, err := db.Query(strQuery, 33, 32)*/

	strQuery := fmt.Sprintf("SELECT * FROM %s.Persons2", *schemaNamePtr)
	//strQuery := fmt.Sprintf("INSERT INTO %s.Persons2 (PersonID, LastName, FirstName, Address, City) VALUES (1, 'LastName1', 'FirstName1', 'Address1', 'City1');", *schemaNamePtr)
	rows, err := db.Query(strQuery)
	if err != nil {
		log.Fatal("Error in Query, Error Info:", err)
	}
	for rows.Next() {
		var dperson, NameLast, FirstName, Address, City string
		if err := rows.Scan(&dperson, &NameLast, &FirstName, &Address, &City); err != nil {
			fmt.Println("Error in Scan, Error Info:", err)
		}
		fmt.Println("PersonID:", dperson, "LastName:", NameLast, "FirstName:", FirstName, "Address:", Address, "City:", City)
	}
	defer rows.Close()

	/*strQuery = fmt.Sprintf("INSERT INTO %s.Persons2 (PersonID, LastName) VALUES (?, ?);", *schemaNamePtr)
	_, err = db.Query(strQuery, 32, "tom")*/

	/*strQuery = fmt.Sprintf("UPDATE %s.Persons2 SET FirstName='cruise', Address='Norwary' WHERE PersonID=32;", *schemaNamePtr)
	_, err = db.Query(strQuery)*/

	/*strQuery = fmt.Sprintf("UPDATE %s.Persons2 SET FirstName=?, Address=? WHERE PersonID=?;", *schemaNamePtr)
	_, err = db.Query(strQuery, "cruise2", "Norwary2", 32)*/

	/*strQuery = fmt.Sprintf("DELETE FROM %s.Persons WHERE PersonID=?;", *schemaNamePtr)
	rows, err = db.Query(strQuery, 34)
	if err != nil {
		log.Fatal("Error in Exec, Error Info:", err)
	}*/

}

//go run m4query.go -host=127.0.0.1 -dbName=postgres -user=postgres -password=postgres -sslmode=disable
//Successfully connected!
