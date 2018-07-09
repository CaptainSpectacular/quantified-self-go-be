package main

import (
    _"github.com/lib/pq"
    "log"
)

type Food struct {
    Id       int    `json:"id"`
    Name     string `json:"name"`
    Calories uint16 `json:"calories"`
}

type Foods []Food


func QueryFoods() []Food{
    // Connect to the Database
    db := ConnectDB()
    defer db.Close()

    // Query for all foods
    rows, err := db.Query("SELECT * FROM foods")
    if err != nil {
        log.Fatal(err)
    }

    // Turn them into food objects, return the foods
    foods := Foods{}

    for rows.Next() {
        food := Food{}
        err := rows.Scan(&food.Id, &food.Name, &food.Calories)
        if err != nil {
            log.Fatal(err)
        }
        foods = append(foods, food)
    }

    return foods
}

