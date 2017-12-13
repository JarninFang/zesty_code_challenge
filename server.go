package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
  fmt.Printf("hi")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
