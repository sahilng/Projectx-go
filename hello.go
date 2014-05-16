package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
    panic(http.ListenAndServe(":17901", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><head></head><body><h1>Welcome Home!</h1></body></html>")
}