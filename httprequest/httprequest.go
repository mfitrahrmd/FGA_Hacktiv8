package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// get request
	res, err := http.Get("https://yesno.wtf/api")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%s", body)
	// end of get request

	// post request
	data := map[string]any{
		"title":  "General Relativity",
		"body":   "the study of time",
		"userId": 1,
	}

	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	client := http.Client{}
	res, err = client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%s", body)
	// end of post request
}
