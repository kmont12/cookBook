package main

import (
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            log.Println(w, "GET, %q", html.EscapeString(r.URL.Path))
        } else if r.Method == "POST" {
            log.Println(w, "POST, %q", html.EscapeString(r.URL.Path))
        } else {
            log.Println(w, "Invalid request method.", 405)
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
