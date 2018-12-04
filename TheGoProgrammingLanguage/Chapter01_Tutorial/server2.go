
package main

import "fmt"
import "log"
import "net/http"
import "sync"

var mu sync.Mutex  // A Mutex is a mutual exclusion lock
var count int

func main() {

    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)

    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

   mu.Lock()
   count++
   mu.Unlock()

   fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {

    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}

// access it through: localhost:8000 or localhost:8000/count
