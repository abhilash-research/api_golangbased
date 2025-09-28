package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "pong")
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello from Go API!")
    })

    port := ":3000"
    fmt.Println("Server listening on port", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}
