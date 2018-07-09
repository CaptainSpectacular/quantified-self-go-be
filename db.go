package main

import (
    "database/sql"
    _"github.com/lib/pq"
    "log"
    "fmt"
)

const (
    host   = ""
    port   = 5432
    user   = "lighthouse"
    dbname = "qs_go"
)

func ConnectDB() *sql.DB {

    dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
                           user, dbname)
    db, err := sql.Open("postgres", dbinfo)

    if err != nil {
        log.Fatal(err)
    }

    return db
}
