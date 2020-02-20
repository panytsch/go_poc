package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"log"
)

type testProcedureResponse struct {
	ReturnCode int
	one        int
	two        string
}

var server = "localhost"
var port = 1433
var user = "sa"
var password = "Qwerty1234"
var database = "master"

func main() {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Got error: %v", err)
	}
	defer db.Close()

	execTestSP(db, 1)
	execTestSP(db, 0)
}

func execTestSP(db *sqlx.DB, v uint8) {
	log.Printf("running sp with parameter %v\n", v)
	//res := &testProcedureResponse{}
	rows, err := db.Queryx(fmt.Sprintf("exec testProcedure %v", v))
	if err != nil {
		log.Fatalf("Error while quering: %v", err)
	}
	rows.Next()
	m := make(map[string]interface{})
	_ = rows.MapScan(m)
	log.Printf("row: %v", m)
	//err = rows.StructScan(res)
	//if err != nil {
	//	log.Fatalf("Error while scanning: %v", err)
	//}
	//log.Printf("Result is: %v\n", res)
}
