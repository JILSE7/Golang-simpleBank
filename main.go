package main

import (
	"database/sql"
	"log"

	"github.com/jilse17/simplebank/api"
	db "github.com/jilse17/simplebank/db/sqlc"

	_ "github.com/lib/pq" // pq to talk with the DB important
)

const (
	dbDvriver     = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	connection, err := sql.Open(dbDvriver, dbSource)

	if err != nil {
		log.Fatal("cannot conncet to db", err)
	}

	store := db.NewStore(connection)

	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("can´t start the server", err)
	}

}
