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
	fmt.Println("DNS:", DNS)
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	/*strQuery := "INSERT INTO Persons2 (PersonID, LastName) VALUES (?, ?);"
	result, err := db.Exec(strQuery, 33, "sagar")*/

	/*strQuery := "UPDATE Persons2 SET FirstName=?, Address=? WHERE PersonID=?;"
	result, err := db.Exec(strQuery, "cruise4", "Norwary4", 33)*/

	/*strQuery := "CREATE TABLE Persons3 (PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255))"
	result, err := db.Exec(strQuery)*/

	strQuery := "DELETE FROM Persons2 WHERE PersonID=?;"
	result, err := db.Exec(strQuery, 32)
	if err != nil {
		log.Fatal("Error in Exec, Error Info:", err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Error in LastInsertId, Error Info:", err)
	}
	fmt.Println("lastInsertId:", lastInsertId)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error in RowsAffected, Error Info:", err)
	}
	fmt.Println("rowsAffected:", rowsAffected)
}
