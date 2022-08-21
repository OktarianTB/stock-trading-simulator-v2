package main

import (
	"database/sql"
	"log"

	"github.com/OktarianTB/stock-trading-simulator-golang/api"
	db "github.com/OktarianTB/stock-trading-simulator-golang/db/sqlc"
	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}