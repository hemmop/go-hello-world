package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    now := time.Now().Format(time.RFC3339)
    fmt.Println("Got request at", now)
    w.Header().Set("X-TimeStamp", now)
    fmt.Fprintln(w, "Hello World @", now)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
