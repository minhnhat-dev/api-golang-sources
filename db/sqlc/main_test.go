package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver =  "postgres"
	dbSource = "postgresql://root:root@localhost/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Connect database error", err)
	}
	
	testQueries = New(testDB)

	os.Exit(m.Run())
}