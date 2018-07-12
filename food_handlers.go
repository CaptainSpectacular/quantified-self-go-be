package main

import (
    "github.com/gorilla/mux"
    "encoding/json"
    "net/http"
    "io/ioutil"
)

type FoodStruct struct {
    Food Food `json:"food"`
}

func FoodIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    foods := QueryFoods()
    json.NewEncoder(w).Encode(foods)
}

func FoodShow(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    id := mux.Vars(r)["id"]
    food := QueryFood(id)
    json.NewEncoder(w).Encode(food)
}

func FoodCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")


    // Unpack Body
    var food FoodStruct 
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &food)

    // Build response and create food
    response := CreateFood(food.Food)
    json.NewEncoder(w).Encode(response)
}

func FoodUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    id := mux.Vars(r)["id"]

    var food FoodStruct
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &food)

    response := UpdateFood(id, food.Food)
    json.NewEncoder(w).Encode(response)
}

func FoodDelete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")


    id := mux.Vars(r)["id"]
    DeleteFood(id)
}
