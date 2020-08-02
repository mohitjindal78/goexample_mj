//Server1 is minimum echo server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)   //each request call handler
	http.HandleFunc("/love", lover) //each request call handler
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//handler echoes the path component of request URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//handler echoes the path component of request URL
func lover(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Toshi Jindal")
}
