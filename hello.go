package main

import(
    "net/http"
)

const resp = `<html>
    <head>
        <title>ProjectX</title>
        <link rel=stylehseet type=text href="style.css">
    </head>
    <body>
        <h1>ProjectXs Dev Site</h1>
        <p>Hello World.</p>
    </body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(resp))
}


func main() {
    http.HandleFunc("/", handler)//homepage
    
    http.Handle("/static", http.FileServer(http.Dir("./static/")))

    http.ListenAndServe(":80", nil)
}
