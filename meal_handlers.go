package main

import (
    "encoding/json"
    "net/http"
)

func MealIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    meals := QueryMeals()
    json.NewEncoder(w).Encode(meals)
}
