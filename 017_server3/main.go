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
	fmt.Fprintf(w, "%v\n", r) //print full request
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//handler echoes the path component of request URL
func lover(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Toshi Jindal")
}