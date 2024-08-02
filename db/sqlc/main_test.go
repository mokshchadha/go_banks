package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/mokshchadha/go_banks/db/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, e := util.LoadConfig("../..")

	if e != nil {
		log.Fatal("Could not load config", e)
	}
	var err error
	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	m.Run()
}
