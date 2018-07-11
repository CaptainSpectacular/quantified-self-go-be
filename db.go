package main

import (
    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/postgres"
    "log"
    "fmt"
    "os"
)


func ConnectDB() *gorm.DB {
    host     := os.Getenv("HOST")
    user     := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    dbname   := os.Getenv("DBNAME")
    sslmode  := os.Getenv("SSLMODE")

    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
                          host, user, password, dbname, sslmode)
    db, err := gorm.Open("postgres", dbinfo)
    if err != nil {
        log.Fatal(err)
    }

    return db
}
