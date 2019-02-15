package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
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

	ctxBackground := context.Background()
	ctx, cancel := context.WithCancel(ctxBackground)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	//_, execErr := tx.ExecContext(ctx, "INSERT INTO Persons2 (PersonID, LastName) VALUES (?, ?);", 4, "cruise4")
	/*if rollbackErr := tx.Rollback(); rollbackErr != nil {
		log.Printf("1 Could not roll back: %v\n", rollbackErr)
	}*/
	_, execErr := tx.ExecContext(ctx, "UPDATE Persons2 SET FirstName=?, Address=? WHERE PersonID=?;", "cruise2", "Norwary2", 33)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("2 Could not roll back: %v\n", rollbackErr)
		}
		log.Println(execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal("tx.Commit:",err)
	}
}
