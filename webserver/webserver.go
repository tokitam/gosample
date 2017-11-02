
package main

import (
  "log"
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is go web server.")
  })

  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}


