package main

import (
    "fmt"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "HomeHandler")
}

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}

func main() {
    http.HandleFunc("/", HomeHandler) // homepage

    // Mandatory root-based resources
    serveSingle("/sitemap.xml", "./sitemap.xml")
    serveSingle("/favicon.ico", "./favicon.ico")
    serveSingle("/robots.txt", "./robots.txt")

    // Normal resources
    http.Handle("/static", http.FileServer(http.Dir("./static/")))

    http.ListenAndServe(":8080", nil)
}