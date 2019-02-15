package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//var ctx = context.Background()
//var db *sql.DB

type key string
const requestPersonIDKey  = key("PersonID")
func main() {
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

	ctxBackground := context.Background()
	ctx, cancel := context.WithCancel(ctxBackground)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()
	// OR
	/*go func() {
		time.Sleep(time.Second)
		cancel() //(send cancel signal to context), log.Print(ctx.Err()) > 2019/02/04 12:32:54 context canceled
	}()*/
	//sleepAndTalk(ctx, 5*time.Second, "hello")

	//time.Sleep(time.Second)
	/*ctx2, cancel2 := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel2()
	ctx = ctx2 // Here WithCancel(ctx) is overwrite with WithTimeout(ctx2). i.e., WithCancel won't send cancelation signal
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("db.PingContext, err:",err)
	}*/
	//time.Sleep(time.Second)

	id := 33
	fmt.Println("requestPersonIDKey:",requestPersonIDKey)
	ctx = context.WithValue(ctx, requestPersonIDKey, id)
	value := getPersonID(ctx)
	strQuery := "UPDATE Persons2 SET FirstName=?, Address=? WHERE PersonID=?;"
	result, err := db.ExecContext(ctx, strQuery, "cruise", "Norwary", value)

	//result, err := db.ExecContext(ctx,strQuery, "cruise2", "Norwary2", id)
	if err != nil {
		log.Fatal("db.ExecContext, err:",err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("result.RowsAffected, err:",err)
	}
	fmt.Println(rowsAffected)
	if rowsAffected != 1 {
		//panic(err)
	}

	time.Sleep(time.Second)
	strQuery = "SELECT * FROM persons"
	rows, err := db.QueryContext(ctx, strQuery, )
	if err != nil {
		log.Fatal("Error in Query, Error Info:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var PersonID, LastName, FirstName, Address, City string
		if err := rows.Scan(&PersonID, &LastName, &FirstName, &Address, &City); err != nil {
			fmt.Println("Error in Scan, Error Info:", err)
		}
		fmt.Println("PersonID:", PersonID, "LastName:", LastName, "FirstName:", FirstName, "Address:", Address, "City:", City)
	}

	strQuery = "SELECT * FROM persons WHERE PersonID=?"
	row := db.QueryRowContext(ctx, strQuery,34)
	var PersonID, LastName, FirstName, Address, City string
	if err := row.Scan(&PersonID, &LastName, &FirstName, &Address, &City); err != nil {
		fmt.Println("Error in Scan, Error Info:", err)
	}
	fmt.Println("PersonID:", PersonID, "LastName:", LastName, "FirstName:", FirstName, "Address:", Address, "City:", City)
}

func getPersonID(ctx context.Context) int{
	id,ok := ctx.Value(requestPersonIDKey).(int)
	if !ok{
		log.Panicln("could not find request ID in context")
	}
	return id
}
func sleepAndTalk(ctx context.Context, duration time.Duration, msg string) {
	select {
	case <-time.After(duration):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print("ctx.Err():",ctx.Err())
	}
}