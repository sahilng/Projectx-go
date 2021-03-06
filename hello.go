package main

import(
    "net/http"
    "log"
    "os"
)

const resp = `<html>
    <head>
        <title>ProjectX</title>
        <link rel=stylehseet type=text href="static/style.css">
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
    http.HandleFunc("/", handler)

    http.Handle("/static", http.FileServer(http.Dir("./static/")))

    err := http.ListenAndServe(":80", nil)

    if err != nil {
        log.Println(err)
        os.Exit(1)
    }
}
