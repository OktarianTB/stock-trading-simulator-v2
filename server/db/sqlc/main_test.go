package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
