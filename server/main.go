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

	log.Println("loaded config successfully")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Printf("driver: %v, source: %v\n", config.DBDriver, config.DBSource)
		log.Fatal("cannot connect to db:", err)
	}
	log.Println("connected to db successfully")

	store := db.NewStore(conn)
	log.Println("created store successfully")

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	log.Println("created server successfully")

	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
