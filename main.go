package main

import (
    "log"
    "database/sql"

    _ "github.com/lib/pq"
    "github.com/JohannSuarez/GoBackend/api"
    "github.com/JohannSuarez/GoBackend/util"
    db "github.com/JohannSuarez/GoBackend/db/sqlc"
)

/*
const (
    dbDriver = "postgres"
    dbSource = "postgresql://root:esth3r@localhost:5432/simple_bank?sslmode=disable"
    serverAddress = "0.0.0.0:8080"
)
*/

func main() {

    config, err := util.LoadConfig(".")

    if err != nil {
        log.Fatal("cannot load config", err)
    }


    conn, err := sql.Open(config.DBDriver, config.DBSource)
    if err != nil {
        log.Fatal("Cannot connect to db", err)

    }

    store := db.NewStore(conn)
    server, err := api.NewServer(config, store)

    if err != nil {
        log.Fatal("Cannot create server:", err)
    }


    err = server.Start(config.ServerAddress)
    if err != nil {
        log.Fatal("Cannot start server:", err)

    }
}
