package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func MealIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    meals := QueryMeals()
    json.NewEncoder(w).Encode(meals)
}

func MealShow(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    id := mux.Vars(r)["id"]
    meal := QueryMeal(id)
    json.NewEncoder(w).Encode(meal)
}

func MealUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    mid := mux.Vars(r)["mid"]
    fid := mux.Vars(r)["fid"]

    message := AddFood(mid, fid)
    json.NewEncoder(w).Encode(message)
}

func MealDelete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    mid := mux.Vars(r)["mid"]
    fid := mux.Vars(r)["fid"]

    message := RemoveFood(mid, fid)
    json.NewEncoder(w).Encode(message)
}
