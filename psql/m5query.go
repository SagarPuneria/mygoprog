package main

import (
	"flag"
	"fmt"

	"github.com/astaxie/beego/orm"
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

type KeyValue struct {
	Personid  int    `json:"emplyeeid"`
	Firstname string `json:"fistname,omitempty"`
	Lastname  string `json:"lastname"`
	Address   string `json:"addrs,omitempty"`
	City      string `json:"city,omitempty"`
}

func main() {

	hostPtr := flag.String("host", host, "a host flag")
	dbNamePtr := flag.String("dbName", dbName, "a dbName flag")
	schemaNamePtr := flag.String("schemaName", schemaName, "a schemaName flag")
	userPtr := flag.String("user", user, "a user flag")
	passwordPtr := flag.String("password", password, "a password flag")
	sslmodePtr := flag.String("sslmode", sslmode, "a user flag")
	flag.Parse()

	DBString := "host=%s port=%d dbname=%s user=%s password='%s' sslmode=%s search_path=%s application_name=employee"
	postgres := fmt.Sprintf(DBString,
		*hostPtr,
		port,
		*dbNamePtr,
		*userPtr,
		*passwordPtr,
		*sslmodePtr,
		*schemaNamePtr)
	fmt.Println("postgres:", postgres)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", postgres)
	o := orm.NewOrm()
	o.Raw("SET SCHEMA employee")
	fmt.Println("Successfully connected!")

	strQuery := fmt.Sprintf("SELECT * FROM %s.Persons2", *schemaNamePtr)

	ormObj := orm.NewOrm()

	var lists []orm.ParamsList
	_, err := ormObj.Raw(strQuery).ValuesList(&lists)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("ValuesList, lists:", lists)

	var keyValue []*KeyValue
	_, err = ormObj.Raw(strQuery).QueryRows(&keyValue)
	if err != nil {
		fmt.Println("err:", err)
	}
	for _, v := range keyValue {
		fmt.Println("QueryRows, value:", *v)
	}
}

/* IN-M-6ZQJG5J:psql sagar.puneria$ go run m5query.go -host=127.0.0.1 -dbName=postgres -user=postgres -password=postgres -sslmode=disable
postgres: host=127.0.0.1 port=5432 dbname=postgres user=postgres password='postgres' sslmode=disable search_path=employee application_name=employee
Successfully connected!
ValuesList, lists: [[2 LastName2 FirstName2 Address2 City2] [2 LastName2 FirstName2 Address2 City2] [1 LastName1 FirstName1 Address1 City1]]
QueryRows, value: {2 FirstName2 LastName2 Address2 City2}
QueryRows, value: {2 FirstName2 LastName2 Address2 City2}
QueryRows, value: {1 FirstName1 LastName1 Address1 City1}
*/
