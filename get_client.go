package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	name := flag.String("name", "", "Name")
	flag.Parse()
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/profile/%s", *name))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
