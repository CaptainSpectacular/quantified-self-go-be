package main

import (
    _"github.com/lib/pq"
    _"github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)

type Meal struct {
    Id    int    `json:"id"`
    Name  string `json:"name"`
    Foods Foods `gorm:"many2many:meal_foods;" json:"foods"`
}

type Message struct {
    M string `json:"message"`
}

type Meals []Meal

func QueryMeals() Meals{
    db := ConnectDB()
    defer db.Close()

    meals := Meals{}
    db.Preload("Foods").Find(&meals)

    return meals
}

func QueryMeal(id string) Meal {
    db := ConnectDB()
    defer db.Close()

    meal := Meal{}
    db.Preload("Foods").First(&meal, id)

    return meal
}

func AddFood(mid string, fid string) Message {
    db := ConnectDB()
    defer db.Close()

    meal    := Meal{}
    food    := Food{}
    message := Message{}

    db.First(&meal, mid)
    db.First(&food, fid)

    meal_food := MealFood{0, food.Id, meal.Id}
    db.NewRecord(meal_food)
    db.Create(&meal_food)

    message.M = fmt.Sprintf("Successfully added %s to %s", food.Name, meal.Name)

    return message
}

func RemoveFood(mid string, fid string) Message {
    db := ConnectDB()
    defer db.Close()

    meal, food, meal_food, message := Meal{}, Food{}, MealFood{}, Message{}
    db.First(&food, fid)
    db.First(&meal, mid)

    message.M = fmt.Sprintf("Successfully removed %s from %s", food.Name, meal.Name)

    db.Where("meal_id = ? AND food_id = ?", meal.Id, food.Id).Find(&meal_food)
    db.Delete(&meal_food)

    return message
}
