package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	user := os.Args[1]
	pass := os.Args[2]
	host := os.Args[3]
	port := os.Args[4]
	DNS := user + ":" + pass + "@tcp(" + host + ":" + port + ")/testDB"
	//DNS := user + ":" + pass + "@tcp(" + host + ":" + port + ")/"
	fmt.Println("DNS:", DNS)
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	db.SetMaxOpenConns(2)

	ctxBackground := context.Background()
	ctx, cancel := context.WithCancel(ctxBackground)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	// A *DB is a pool of connections. Call Conn to reserve a connection for
	// exclusive use.
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal("conn ctx err:", err)
	}
	defer conn.Close() // Return the connection to the pool.
	strQuery := "INSERT INTO Persons VALUES(35, 'tom5', 'cruise5', 'Norwary5', 'Address5')"
	stmt, err := conn.PrepareContext(ctx, strQuery)
	if err != nil {
		log.Fatal("Error in Prepare, Error Info:", err)
	}
	result, err := stmt.ExecContext(ctx)
	if err != nil {
		log.Fatal("db.ExecContext, err:",err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("result.RowsAffected, err:",err)
	}
	fmt.Println("debug rowsAffected:",rowsAffected)

	conn1, err := db.Conn(ctx)
	if err != nil {
		log.Fatal("conn1 ctx err:", err)
	}
	defer conn1.Close() // Return the connection to the pool.
	strQuery = "INSERT INTO Persons VALUES(36, 'tom6', 'cruise6', 'Norwary6', 'Address6')"
	result, err = conn1.ExecContext(ctx, strQuery)
	if err != nil {
		log.Fatal("db.ExecContext, err:",err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Fatal("result.RowsAffected, err:",err)
	}
	fmt.Println("debug 1 rowsAffected:",rowsAffected)

	fmt.Println("waiting at db.Conn. Because of db.SetMaxOpenConns(2) OR close atleast one connection(conn.Close or conn1.Close) to get new conn2")
	conn2, err := db.Conn(ctx)
	if err != nil {
		log.Fatal("conn2 ctx err:", err)
	}
	fmt.Println("debug 2 conn2")
	defer conn2.Close() // Return the connection to the pool.
	strQuery = "INSERT INTO Persons VALUES(37, 'tom7', 'cruise7', 'Norwary7', 'Address7')"
	result, err = conn2.ExecContext(ctx, strQuery)
	if err != nil {
		log.Fatal("db.ExecContext, err:",err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Fatal("result.RowsAffected, err:",err)
	}
	fmt.Println(rowsAffected)
}