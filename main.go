package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // without this the server and db cannot communicate
	"github.com/mokshchadha/go_banks/api"
	db "github.com/mokshchadha/go_banks/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://moksh:@localhost:5432/simple_banks?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to the db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
