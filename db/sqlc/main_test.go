package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://postgres:postgres@localhost:5432/receita_na_mao?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())

}
