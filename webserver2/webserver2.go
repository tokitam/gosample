
package main

import (
  "log"
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
)

func Exists(name string) bool {
   _, err := os.Stat(name)
  return !os.IsNotExist(err)
}
	
func main() {
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
  })

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //var path = "/var/www/html"
    var path = "/Library/WebServer/Documents"
    var f = path + r.URL.Path
    
    if (!Exists(f)) {
      fmt.Println("File not found: " + f);
      return
    }

    dat, err := ioutil.ReadFile(f)
    if err != nil {
      panic(err)
    }

    fmt.Fprint(w, string(dat))
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}


