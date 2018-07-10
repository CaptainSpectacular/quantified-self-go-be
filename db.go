package main

import (
    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/postgres"
    "log"
    "fmt"
)

const (
    host   = ""
    port   = 5432
    user   = "lighthouse"
    dbname = "qs_go"
)

func ConnectDB() *gorm.DB {
    dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
                           user, dbname)
    db, err := gorm.Open("postgres", dbinfo)
    if err != nil {
        log.Fatal(err)
    }

    return db
}
