package main

import(
    "log"
    "net/http"
    "os"
    "github.com/rs/cors"
)

func main() {
    router := NewRouter()
    c := cors.New(cors.Options{
        AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
        AllowedHeaders: []string{"Access-Control-Allow-Origin"},
    })
    handler := c.Handler(router) 
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(http.ListenAndServe(":"+port, handler))
}


