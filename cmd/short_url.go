package main

import (
	"github.com/dannylee1020/url-shortener/internal/data"
	"net/http"
)

func (app *application) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ShortURL string `json:"short_url,omitempty"`
		LongURL  string `json:"long_url,omitempty"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.invalidRequestResponse(w, r, err)
	}

	urlData := &data.UrlData{
		LongURL: input.LongURL,
	}

	_, err = app.models.URL.QueryWithLong(input.LongURL)

	if err == nil {
		message := "shortened url already exists in the DB"
		app.errorResponse(w, r, http.StatusConflict, message)
		return
	}

	shortUrl := data.GenerateShortURL()
	err = app.models.URL.InsertURL(urlData, shortUrl)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	urlData.ShortURL = shortUrl

	err = app.writeJSON(w, http.StatusOK, urlData, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
