package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello/:name", Hello)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
