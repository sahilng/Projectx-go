package main

import(
    "net/http"
    "log"
    "os"
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

func cssHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "style.css")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/style.css", cssHandler)
    err := http.ListenAndServe(":80", nil)

    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    if err2 != nil {
        log.Println(err2)
        os.Exit(1)
    }
}
