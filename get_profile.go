package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Profile struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	FavoriteFoods []string `json:"favorite_foods"`
}

var bob Profile = Profile{
	Name:          "Bob",
	Age:           25,
	Gender:        "Man",
	FavoriteFoods: []string{"Hamburger", "Cookie", "Chocolate"},
}

var alice Profile = Profile{
	Name:          "Alice",
	Age:           24,
	Gender:        "Woman",
	FavoriteFoods: []string{"Apple", "Orange", "Melon"},
}

func GetProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	var resProfile Profile
	if name == "Bob" {
		resProfile = bob
	} else if name == "Alice" {
		resProfile = alice
	} else {
		http.Error(w, fmt.Sprintf("%d Not Found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(resProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintf(w, string(bytes))
}

func main() {
	router := httprouter.New()
	router.GET("/profile/:name", GetProfile)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
