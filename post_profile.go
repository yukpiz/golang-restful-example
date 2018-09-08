package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

type Profile struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	FavoriteFoods []string `json:"favorite_foods"`
}

var savedProfiles map[string]Profile = map[string]Profile{}

func PostProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var reqProfile Profile
	err = json.Unmarshal(body, &reqProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqProfile.Name == "" {
		http.Error(w, "name is required.", http.StatusBadRequest)
		return
	}

	savedProfiles[reqProfile.Name] = reqProfile
	fmt.Fprintf(w, fmt.Sprintf("%d Created", http.StatusCreated))
}

func main() {
	router := httprouter.New()
	router.POST("/profile", PostProfile)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
