
package main

import "fmt"
import "net/http"
import "log"

// To check the details of URL, please visit: https://golang.org/pkg/net/url/#URL

func main() {
    // HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
    http.HandleFunc("/", handler)

    // ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
