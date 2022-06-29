package db

import (
    "os"
    "log"
    "database/sql"
    "testing"
    "github.com/JohannSuarez/GoBackend/util"
    _ "github.com/lib/pq"
)

/*
const (
    dbDriver = "postgres"
    dbSource = "postgresql://root:esth3r@localhost:5432/simple_bank?sslmode=disable"
)
*/

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
    var err error
    config, err := util.LoadConfig("../../")
    // testDB, err = sql.Open(dbDriver, dbSource)
    testDB, err = sql.Open(config.DBDriver, config.DBSource)

    if err != nil {
        log.Fatal("Cannot connect to the database: ", err)
    }

    testQueries = New(testDB)

    os.Exit(m.Run())
}
