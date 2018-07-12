package main

import(
    "log"
    "net/http"
    "os"
)

func main() {
    router := NewRouter()
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(http.ListenAndServe(":"+port, router))
}

func EnableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

