package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Profile struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	FavoriteFoods []string `json:"favorite_foods"`
}

func main() {
	profile := parseArgs()
	fmt.Printf("%+v\n", profile)

	jsonStr, err := json.Marshal(profile)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/profile", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func parseArgs() Profile {
	name := flag.String("name", "", "Name")
	age := flag.Int("age", 0, "Age")
	gender := flag.String("gender", "", "Gender")
	foods := flag.String("foods", "", "Favorite foods")
	flag.Parse()

	return Profile{
		Name:          *name,
		Age:           *age,
		Gender:        *gender,
		FavoriteFoods: strings.Split(*foods, " "),
	}
}
