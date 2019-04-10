package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/employeedb")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	go func() {
		for {
			err := db.Ping()
			if err != nil {
				log.Fatal("Error in ping, Error Info:", err)
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Println("After 500 Millisecond sleep")
		}
	}()
	strQuery := "SELECT * FROM persons2"
	rows, err := db.Query(strQuery)
	if err != nil {
		log.Fatal("Error in Query, Error Info:", err)
	}
	columnType, err := rows.ColumnTypes()
	if err !=nil {
		fmt.Println("Error in ColumnTypes, Error Info:", err)
	}
	for _, c := range columnType {
		fmt.Println(">Name:", c.Name())
		fmt.Println("DatabaseTypeName:", c.DatabaseTypeName())
		precision, scale, ok := c.DecimalSize()
		fmt.Println("precision:", precision, "scale:", scale, "ok:",ok)
		nullable, ok := c.Nullable()
		fmt.Println("nullable:", nullable, ",ok:", ok)
		fmt.Println("ScanType:", c.ScanType())
	}
	fmt.Println("===================================")
	rowHeading, _ := rows.Columns()
	fmt.Println("rowHeading:", rowHeading)
	ColumnsCount := len(rowHeading)
	vals := make([]interface{}, ColumnsCount)
	args := make([]sql.NullString, ColumnsCount)
	for i, _ := range args {
		vals[i] = &args[i]
	}
	fmt.Println("===================================")
	for rows.Next() {
		if err := rows.Scan(vals...); err != nil {
			fmt.Println("rows.Scan error info: ", err)
		}
		fmt.Println("3args:", args)
		for _,argument := range args{
			if argument.Valid {
				fmt.Println(argument.String)
			}
		}
	}
	time.Sleep(1 * time.Second)
}