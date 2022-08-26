package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vinicius-gregorio/receita_na_mao/api"
	db "github.com/vinicius-gregorio/receita_na_mao/db/sqlc"
	"github.com/vinicius-gregorio/receita_na_mao/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
