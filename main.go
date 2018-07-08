package main

import("fmt"
       //"encoding/json"
       "database/sql"
       _ "github.com/lib/pq"
       "log"
       "net/http"
       "github.com/gorilla/mux")

const (
    host     = "localhost"
    port     = 5432
    user     = "lighthouse"
    dbname   = "qs_go"
)

type Food struct {
    id       int
    name     string
    calories uint16
}


func main() {


    router := mux.NewRouter()
    router.HandleFunc("/foods", GetFoods).Methods("GET")
    router.HandleFunc("/foods", CreateFood).Methods("POST")
    router.HandleFunc("/foods/{id}", GetFood).Methods("GET")
    router.HandleFunc("/foods/{id}", DeleteFood).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":3000", router))
}

func GetFoods(w http.ResponseWriter, r *http.Request) {
    dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
                           user, dbname)
    db, err := sql.Open("postgres", dbinfo)

    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM foods")

    fmt.Println(rows)


    foods := make([]Food, 0)

    for rows.Next() {
        food := Food{}
        err := rows.Scan(&food.id, &food.name, &food.calories)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(food.id, food.name, food.calories)
        foods = append(foods, food)
    }
    fmt.Fprint(w, foods)
}
func GetFood(w http.ResponseWriter, r *http.Request) {}
func CreateFood(w http.ResponseWriter, r *http.Request) {}
func DeleteFood(w http.ResponseWriter, r *http.Request) {}
