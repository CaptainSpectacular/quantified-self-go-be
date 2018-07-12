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

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
