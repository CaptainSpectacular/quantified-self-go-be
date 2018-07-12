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
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
    })
    handler := c.Handler(router)
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(http.ListenAndServe(":"+port, handler))
}


