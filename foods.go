package main

import (
    _"github.com/lib/pq"
    _"github.com/jinzhu/gorm/dialects/postgres"
)

type Food struct {
    Id       int    `json:"id"`
    Name     string `json:"name"`
    Calories int    `json:"calories"`
}

type Foods []Food


func QueryFoods() []Food{
    // Connect to the Database
    db := ConnectDB()
    defer db.Close()

    foods := Foods{}
    db.Find(&foods)

    return foods
}

func QueryFood(id string) Food {
    // Connect to db
    db := ConnectDB()
    defer db.Close()

    food := Food{}
    db.First(&food, id)

    return food
}

func CreateFood(food Food) Food {
    db := ConnectDB()
    defer db.Close()

    db.NewRecord(food)
    db.Create(&food)
    return food
}

func UpdateFood(id string, params Food) Food {
    db := ConnectDB()
    defer db.Close()

    food := Food{}
    db.First(&food, id)
    food.Name     = params.Name
    food.Calories = params.Calories
    db.Save(&food)
    return food
}

func DeleteFood(id string) {
    db := ConnectDB()
    defer db.Close()

    food := Food{}
    db.First(&food, id)
    db.Delete(&food)
}
