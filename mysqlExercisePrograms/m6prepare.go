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

	strQuery := "UPDATE Persons2 SET FirstName=?, Address=? WHERE PersonID=?;"
	stmt, err := db.Prepare(strQuery)
	if err != nil {
		log.Fatal("Error in Prepare, Error Info:", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec("cruise3", "Norwary3", 33)
	if err != nil {
		log.Fatal("Error in Exec, Error Info:", err)
	}

	/*strQuery := "INSERT INTO Persons2 (PersonID, LastName) VALUES (?, ?);"
	stmt, err := db.Prepare(strQuery)
	if err != nil {
		log.Fatal("Error in Prepare, Error Info:", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(34, "tom")
	if err != nil {
		log.Fatal("Error in Exec, Error Info:", err)
	}*/

	/*strQuery := "CREATE TABLE Persons4 (PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255))"
	stmt, err := db.Prepare(strQuery)
	if err != nil {
		log.Fatal("Error in Prepare, Error Info:", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec()
	if err != nil {
		log.Fatal("Error in Exec, Error Info:", err)
	}*/

	/*strQuery := "DELETE FROM Persons2 WHERE PersonID=?;"
	stmt, err := db.Prepare(strQuery)
	if err != nil {
		log.Fatal("Error in Prepare, Error Info:", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(34)
	if err != nil {
		log.Fatal("Error in Exec, Error Info:", err)
	}*/

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
