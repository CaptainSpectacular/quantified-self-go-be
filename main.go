package main

import(
    "log"
    "net/http"
    "github.com/gorilla/handlers"
    "os"
)

func main() {
    router := NewRouter()
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
    allowedOrigins := handlers.AllowedOrigins([]string{"*"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
    log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}

