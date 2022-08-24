package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vinicius-gregorio/receita_na_mao/api"
	db "github.com/vinicius-gregorio/receita_na_mao/db/sqlc"
)

const (
	DBDriver      = "postgres"
	DBSource      = "postgresql://postgres:postgres@localhost:5432/receita_na_mao?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(serverAddress)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
