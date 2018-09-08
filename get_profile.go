package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func GetProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func main() {
	router := httprouter.New()
	router.GET("/profile/:name", GetProfile)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
