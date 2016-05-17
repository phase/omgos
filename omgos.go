package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("GET %s\n", r.URL.Path[1:])
    data, err := ioutil.ReadFile(r.URL.Path[1:])
    if err != nil {
        fmt.Fprintf(w, "%q\n", err)
    } else {
        w.Write(data)
    }
}

func main() {
    fmt.Println("Starting HTTP Server")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}