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

    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
                          host, user, password, dbname)
    db, err := gorm.Open("postgres", dbinfo)
    if err != nil {
        log.Fatal(err)
    }

    return db
}
