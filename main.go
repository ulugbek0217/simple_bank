package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ulugbek0217/simple_bank/api"
	db "github.com/ulugbek0217/simple_bank/db/sqlc"
	"github.com/ulugbek0217/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load env:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("error starting server:", err)
	}
}
