package main

import (
    _"github.com/lib/pq"
    _"github.com/jinzhu/gorm/dialects/postgres"
)

type Meal struct {
    Id    int    `json:"id"`
    Name  string `json:"name"`
    Foods Foods `gorm:"many2many:meal_foods;" json:"foods"`
}

type Meals []Meal

func QueryMeals() Meals{
    db := ConnectDB()
    defer db.Close()

    meals := Meals{}
    db.Preload("Foods").Find(&meals)

    return meals
}
