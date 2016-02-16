package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("request: method=", r.Method, " URL=", r.URL, " Proto=", r.Proto, " Remote=", r.RemoteAddr)
    fmt.Fprintf(w, "Hello World! %s", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
