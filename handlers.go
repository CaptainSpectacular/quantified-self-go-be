package main

import (
    "github.com/gorilla/mux"
    "encoding/json"
    "net/http"
)


func FoodIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    foods := QueryFoods()
    json.NewEncoder(w).Encode(foods)
}

func FoodShow(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    id := mux.Vars(r)["id"]
    food := QueryFood(id)
    json.NewEncoder(w).Encode(food)
}
func CreateFood(w http.ResponseWriter, r *http.Request) {}
func DeleteFood(w http.ResponseWriter, r *http.Request) {}
