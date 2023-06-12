package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var input struct {
		LongURL string `json:"longURL"`
	}

	var output struct {
		URL string `json:"long-url"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
	}

	output.URL = input.LongURL

	res, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
