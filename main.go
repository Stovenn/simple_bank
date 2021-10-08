package main

import (
	"github.com/stovenn/simple_bank/db/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/stovenn/simple_bank/api"
	db "github.com/stovenn/simple_bank/db/sqlc"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("error while loading config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("error while connecting to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("error while starting server: ", err)
	}
}
