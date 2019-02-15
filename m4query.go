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
	/*strQuery := `CREATE TABLE Persons2 (PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255));
	INSERT INTO Persons2 (PersonID) VALUES (?);`
	rows, err := db.Query(strQuery, 32)*/

	/*strQuery := "CREATE TABLE Persons2 (PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255))"
	rows, err := db.Query(strQuery)*/

	/*strQuery := "SELECT * FROM persons WHERE PersonID=?"
	rows, err := db.Query(strQuery, 33)*/

	strQuery := "SELECT * FROM persons WHERE PersonID=? OR PersonID=?"
	rows, err := db.Query(strQuery, 33, 32)
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
	defer rows.Close()

	/*strQuery = "INSERT INTO Persons2 (PersonID, LastName) VALUES (?, ?);"
	_, err = db.Query(strQuery, 32, "tom")*/

	/*strQuery = "UPDATE Persons2 SET FirstName='cruise', Address='Norwary' WHERE PersonID=32;"
	_, err = db.Query(strQuery)*/

	/*strQuery = "UPDATE Persons2 SET FirstName=?, Address=? WHERE PersonID=?;"
	_, err = db.Query(strQuery, "cruise2", "Norwary2", 32)*/

	strQuery = "DELETE FROM Persons WHERE PersonID=?;"
	rows, err = db.Query(strQuery, 34)
	if err != nil {
		log.Fatal("Error in Exec, Error Info:", err)
	}

}
